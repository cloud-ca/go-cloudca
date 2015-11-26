package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type CCAError struct {
	Code int
	Message string
	Context map[string]interface{}
}

type CCAErrors []CCAError

func (ccaErrors CCAErrors) Error() string {
	return "Number of api errors: " + strconv.Itoa(len(ccaErrors))
}

type CCAResponse struct {
	StatusCode int
	Data []byte
	Errors []CCAError
	MetaData map[string]interface{}
}

func buildErrors(errorResponse []interface{}) []CCAError {
	errors := []CCAError{}
	for _, val := range errorResponse {
		errorMap := val.(map[string]interface{})
		code, _ := errorMap["code"].(int)
		message, _ := errorMap["message"].(string)
		context, _ := errorMap["context"].(map[string]interface{})
		errors = append(errors, CCAError{code, message, context})
	}
	return errors
}

func NewCCAResponse(response *http.Response) (*CCAResponse, error) {
	respBody, err := ioutil.ReadAll(response.Body)
    if err != nil {
		return nil, err
	}
	ccaResponse := CCAResponse{}
	ccaResponse.StatusCode = response.StatusCode
	responseMap := map[string]*json.RawMessage{}
	json.Unmarshal(respBody, &responseMap)

	if val, ok := responseMap["data"]; ok {
		ccaResponse.Data = []byte(*val)
	}

	if val, ok := responseMap["metadata"]; ok {
		metadata := map[string]interface{}{}
		json.Unmarshal(*val, &metadata)
		ccaResponse.MetaData = metadata
	}

	if val, ok := responseMap["errors"]; ok {
		errors := []interface{}{}
		json.Unmarshal(*val, &errors)
		ccaResponse.Errors = buildErrors(errors)
	} else if(response.StatusCode != 200) {
		//should always have errors in response body if not 200 OK
		panic("Unexpected. Received status " + response.Status + " but no errors in response body")
	}
	return &ccaResponse, nil
}