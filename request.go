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

func (ccaReq CCARequest) buildUrl(endpoint string, parameters map[string]string) string  {
	query := url.Values{}
	for k, v := range parameters {
		query.Add(k, v)
	}
	u, _ := url.Parse(ccaReq.apiURL + "/" +  strings.Trim(endpoint, "/") + "?" + query.Encode())
	return u.String()
}

func (ccaReq CCARequest) execute(method string, endpoint string, parameters map[string]string, body []byte) (*CCAResponse, error) {
	client := &http.Client{}
	var bodyBuffer io.Reader
	if body != nil {
		bodyBuffer = bytes.NewBuffer(body)
	}
	req, err := http.NewRequest(method, ccaReq.buildUrl(endpoint, parameters), bodyBuffer)
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

func (ccaReq CCARequest) Get(endpoint string, parameters map[string]string) (*CCAResponse, error) {
	return ccaReq.execute("GET", endpoint, parameters, nil)
}

func (ccaReq CCARequest) Post(endpoint string, parameters map[string]string, body []byte) (*CCAResponse, error) {
	return ccaReq.execute("POST", endpoint, parameters, body)
}

func (ccaReq CCARequest) Del(endpoint string, parameters map[string]string, body []byte) (*CCAResponse, error) {
	return ccaReq.execute("DELETE", endpoint, parameters, body)
}

func (ccaReq CCARequest) Put(endpoint string, parameters map[string]string, body []byte) (*CCAResponse, error) {
	return ccaReq.execute("PUT", endpoint, parameters, body)
}