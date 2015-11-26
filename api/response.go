package api

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
) 

const (
	OK = 200
)

type CcaError struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Context map[string]interface{} `json:"context"`
}

type CcaResponse struct {
	TaskId string
	TaskStatus string
	StatusCode int
	Data []byte
	Errors []CcaError
	MetaData map[string]interface{}
}

type CcaErrorResponse CcaResponse

func (errorResponse CcaErrorResponse) Error() string {
	var errorStr string = "[ERROR] Received HTTP status code " + strconv.Itoa(errorResponse.StatusCode) + "\n"
	for _, e := range errorResponse.Errors {
		context, _ := json.Marshal(e.Context)
		errorStr += "[ERROR] Error Code: " + strconv.Itoa(e.Code) + ", Message: " + e.Message + ", Context: " + string(context)
	}
	return errorStr
}

func (ccaResponse CcaResponse) IsError() bool {
	return ccaResponse.StatusCode != OK
}

func NewCcaResponse(response *http.Response) (CcaResponse, error) {
	respBody, err := ioutil.ReadAll(response.Body)
    if err != nil {
		return CcaResponse{}, err
	}
	ccaResponse := CcaResponse{}
	ccaResponse.StatusCode = response.StatusCode
	responseMap := map[string]*json.RawMessage{}
	json.Unmarshal(respBody, &responseMap)

	if val, ok := responseMap["taskId"]; ok {
		ccaResponse.TaskId = string(*val)
	}

	if val, ok := responseMap["taskStatus"]; ok {
		ccaResponse.TaskStatus = string(*val)
	}

	if val, ok := responseMap["data"]; ok {
		ccaResponse.Data = []byte(*val)
	}

	if val, ok := responseMap["metadata"]; ok {
		metadata := map[string]interface{}{}
		json.Unmarshal(*val, &metadata)
		ccaResponse.MetaData = metadata
	}

	if val, ok := responseMap["errors"]; ok {
		errors := []CcaError{}
		json.Unmarshal(*val, &errors)
		ccaResponse.Errors = errors
	} else if(response.StatusCode != OK) {
		//should always have errors in response body if not 200 OK
		panic("Unexpected. Received status " + response.Status + " but no errors in response body")
	}
	return ccaResponse, nil
}