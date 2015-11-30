package api

import (
	"net/url"
	"bytes"
	"strings"
	"net/http"
	"io"
	)

type CcaClient interface {
	Do(request CcaRequest) (*CcaResponse, error)
}

type CcaApiClient struct {
	apiURL string
	apiKey string
	httpClient *http.Client
}

const API_KEY_HEADER = "MC-Api-Key"

func NewApiClient(apiURL, apiKey string) CcaApiClient {
	return CcaApiClient{
		apiURL: apiURL,
		apiKey: apiKey,
		httpClient: &http.Client{},
	}
}

//Build a URL by using endpoint and options. Options will be set as query parameters.
func (ccaClient CcaApiClient) buildUrl(endpoint string, options map[string]string) string  {
	query := url.Values{}
	if options != nil {
		for k, v := range options {
			query.Add(k, v)
		}
	}
	u, _ := url.Parse(ccaClient.apiURL + "/" +  strings.Trim(endpoint, "/") + "?" + query.Encode())
	return u.String()
}

//Does the API call to server and returns a CCAResponse. Cloud.ca errors will be returned in the
//CCAResponse body, not in the error return value. The error return value is reserved for unexpected errors.
func (ccaClient CcaApiClient) Do(request CcaRequest) (*CcaResponse, error) {
	var bodyBuffer io.Reader
	if request.Body != nil {
		bodyBuffer = bytes.NewBuffer(request.Body)
	}
	method := request.Method
	if method == "" {
		method = "GET"
	}
	req, err := http.NewRequest(request.Method, ccaClient.buildUrl(request.Endpoint, request.Options), bodyBuffer)
	if err != nil {
		return nil, err
	}
	req.Header.Add(API_KEY_HEADER, ccaClient.apiKey)
	resp, err := ccaClient.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return NewCcaResponse(resp)
}