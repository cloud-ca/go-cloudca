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

//Create a CcaClient with the default URL
func NewCcaClient(apiKey string) *CcaClient {
	return NewCcaClientWithURL(DEFAULT_API_URL, apiKey)
}

//Create a CcaClient with a custom URL
func NewCcaClientWithURL(apiURL string, apiKey string) *CcaClient {
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
//For now it assumes that the serviceCode belongs to a cloud.ca service type
func (c CcaClient) GetResources(serviceCode string, environmentName string) services.ServiceResources {
	//TODO: change to check service type of service code
	return cloudca.NewResources(c.apiClient, serviceCode, environmentName)
}

//Get the API url used to do he calls
func (c CcaClient) GetApiURL() string {
	return c.apiURL
}

//Get the API key used in the calls
func (c CcaClient) GetApiKey() string {
	return c.apiKey
}

//Get the API Client used by all the services
func (c CcaClient) GetApiClient() api.CcaApiClient {
	return c.apiClient
}
