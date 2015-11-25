package main

import "fmt"

type CCAClient struct {
	ccaRequest CCARequest
	Tasks TaskApi
}

func NewCCAClient(apiURL string, apiKey string) CCAClient {
	ccaClient := CCAClient{}
	ccaClient.ccaRequest = CCARequest{apiURL, apiKey}
	ccaClient.Tasks = TaskApi{ccaClient.ccaRequest}
	return ccaClient
}

func (c CCAClient) GetResourceApi(serviceCode string, environmentName string) {

}
