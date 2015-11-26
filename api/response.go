package api

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type CCAError struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Context map[string]interface{} `json:"context"`
}

type CCAErrors []CCAError

//TODO: change better error message
func (ccaErrors CCAErrors) Error() string {
	return "Number of api errors: " + strconv.Itoa(len(ccaErrors))
}

type CCAResponse struct {
	TaskId string
	TaskStatus string
	StatusCode int
	Data []byte
	Errors []CCAError
	MetaData map[string]interface{}
}

func NewCCAResponse(response *http.Response) (CCAResponse, error) {
	respBody, err := ioutil.ReadAll(response.Body)
    if err != nil {
		return CCAResponse{}, err
	}
	ccaResponse := CCAResponse{}
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
		errors := []CCAError{}
		json.Unmarshal(*val, &errors)
		ccaResponse.Errors = errors
	} else if(response.StatusCode != 200) {
		//should always have errors in response body if not 200 OK
		panic("Unexpected. Received status " + response.Status + " but no errors in response body")
	}
	return ccaResponse, nil
}