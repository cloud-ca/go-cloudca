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
	Data interface{}
	Errors []CCAError
	MetaData map[string]interface{}
}

//Build CCAError objects from an error field in api response
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

//Build a CCAResponse from an api response
func NewCCAResponse(response *http.Response) (*CCAResponse, error) {
	respBody, err := ioutil.ReadAll(response.Body)
    if err != nil {
		return nil, err
	}
	ccaResponse := CCAResponse{}
	ccaResponse.StatusCode = response.StatusCode
	responseMap := map[string]interface{}{}
	json.Unmarshal(respBody, &responseMap)

	if val, ok := responseMap["data"]; ok {
		ccaResponse.Data = val
	}

	if val, ok := responseMap["metadata"]; ok {
		ccaResponse.MetaData = val.(map[string]interface{})
	}

	if val, ok := responseMap["errors"]; ok {
		ccaResponse.Errors = buildErrors(val.([]interface{}))
	} else if(response.StatusCode != 200) {
		//should always have errors in response body if not 200 OK
		panic("Unexpected. Received status " + response.Status + " but no errors in response body")
	}
	return &ccaResponse, nil
}