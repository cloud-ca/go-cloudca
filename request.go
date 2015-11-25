package main

import (
	"net/url"
	"bytes"
	"strings"
	"net/http"
	"io"
	)

const API_KEY_HEADER = "MC-Api-Key"

type CCARequest struct {
	apiURL string
	apiKey string
}

//Build a URL by using endpoint and options. Options will be set as query parameters.
func (ccaReq CCARequest) buildUrl(endpoint string, options map[string]string) string  {
	query := url.Values{}
	for k, v := range options {
		query.Add(k, v)
	}
	u, _ := url.Parse(ccaReq.apiURL + "/" +  strings.Trim(endpoint, "/") + "?" + query.Encode())
	return u.String()
}

//Does the API call to server and returns a CCAResponse. Cloud.ca errors will be returned in the
//CCAResponse body, not in the error return value. The error return value is reserved for unexpected errors.
func (ccaReq CCARequest) call(method string, endpoint string, options map[string]string, body []byte) (*CCAResponse, error) {
	client := &http.Client{}
	var bodyBuffer io.Reader
	if body != nil {
		bodyBuffer = bytes.NewBuffer(body)
	}
	req, err := http.NewRequest(method, ccaReq.buildUrl(endpoint, options), bodyBuffer)
	if err != nil {
		return nil, err
	}
	req.Header.Add(API_KEY_HEADER, ccaReq.apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return NewCCAResponse(resp)
}

//Does a GET and returns a CCAResponse. Cloud.ca errors will be returned in the
//CCAResponse body, not in the error return value. The error return value is reserved for unexpected errors.
func (ccaReq CCARequest) Get(endpoint string, options map[string]string) (*CCAResponse, error) {
	return ccaReq.call("GET", endpoint, options, nil)
}

//Does a POST and returns a CCAResponse. Cloud.ca errors will be returned in the
//CCAResponse body, not in the error return value. The error return value is reserved for unexpected errors.
func (ccaReq CCARequest) Post(endpoint string, options map[string]string, body []byte) (*CCAResponse, error) {
	return ccaReq.call("POST", endpoint, options, body)
}

//Does a DELETE and returns a CCAResponse. Cloud.ca errors will be returned in the
//CCAResponse body, not in the error return value. The error return value is reserved for unexpected errors.
func (ccaReq CCARequest) Del(endpoint string, options map[string]string, body []byte) (*CCAResponse, error) {
	return ccaReq.call("DELETE", endpoint, options, body)
}

//Does a PUT and returns a CCAResponse. Cloud.ca errors will be returned in the
//CCAResponse body, not in the error return value. The error return value is reserved for unexpected errors.
func (ccaReq CCARequest) Put(endpoint string, options map[string]string, body []byte) (*CCAResponse, error) {
	return ccaReq.call("PUT", endpoint, options, body)
}