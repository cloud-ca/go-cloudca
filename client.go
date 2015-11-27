package gocca

import (
	"github.com/cloud-ca/go-cloudca/services/cloudca"
	"github.com/cloud-ca/go-cloudca/services"
	"github.com/cloud-ca/go-cloudca/api"
)

const (
	DEFAULT_API_URL = "https://api.cloud.ca/v1/"
)

type CcaClient struct {
	apiURL string
	apiKey string
	apiClient api.CcaApiClient
	Tasks services.TaskService
}

func NewCcaClient(apiKey string) *CcaClient {
	return NewCCAClientWithCustomURL(DEFAULT_API_URL, apiKey)
}

func NewCcaClientWithCustomURL(apiURL string, apiKey string) *CcaClient {
	apiClient := api.NewApiClient(apiURL, apiKey)
	ccaClient := CcaClient{
		apiURL: apiURL,
		apiKey: apiKey,
		apiClient: apiClient,
		Tasks: services.NewTaskService(apiClient),
	}
	return &ccaClient
}

//Get the Resources for a specific serviceCode and environmentName
//For now it assumes that the serviceCode belongs to a cloud.ca service
func (c CcaClient) GetResources(serviceCode string, environmentName string) services.ServiceResources {
	//TODO: change to check service type of service code
	return cloudca.NewResources(c.apiClient, serviceCode, environmentName)
}

func (c CcaClient) GetApiURL() string {
	return c.apiURL
}

func (c CcaClient) GetApiKey() string {
	return c.apiKey
}


func (c CcaClient) GetApiClient() api.CcaApiClient {
	return c.apiClient
}
