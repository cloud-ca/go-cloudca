package gocca

import (
	"github.com/cloud-ca/go-cloudca/services/cloudca"
	"github.com/cloud-ca/go-cloudca/services"
	"github.com/cloud-ca/go-cloudca/api"
)

const (
	DEFAULT_API_URL = "https://api.cloud.ca/v1/"
)

type CCAClient struct {
	apiURL string
	apiKey string
	apiClient api.CCAApiClient
	Tasks services.TaskService
}

func NewCCAClient(apiKey string) CCAClient {
	return NewCCAClientWithCustomURL(DEFAULT_API_URL, apiKey)
}

func NewCCAClientWithCustomURL(apiURL string, apiKey string) CCAClient {
	apiClient := api.NewApiClient(apiURL, apiKey)
	ccaClient := CCAClient{
		apiURL: apiURL,
		apiKey: apiKey,
		apiClient: apiClient,
		Tasks: services.NewTaskService(apiClient),
	}
	return ccaClient
}

//Get the Resources for a specific serviceCode and environmentName
//For now it assumes that the serviceCode belongs to a cloud.ca service
func (c CCAClient) GetResources(serviceCode string, environmentName string) interface{} {
	//TODO: change to check service type of service code
	return cloudca.NewResources(c.apiClient, serviceCode, environmentName)
}

func (c CCAClient) GetApiURL() string {
	return c.apiURL
}

func (c CCAClient) GetApiKey() string {
	return c.apiKey
}


func (c CCAClient) GetApiClient() api.CCAApiClient {
	return c.apiClient
}