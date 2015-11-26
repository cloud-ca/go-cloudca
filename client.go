package main

const (
	DEFAULT_API_URL = "https://api.cloud.ca/v1/"
)

type CCAClient struct {
	apiClient CCAApiClient
	Tasks TaskService
}

func NewCCAClient(apiKey string) CCAClient {
	ccaClient := CCAClient{}
	ccaClient.apiClient = CCAApiClient{DEFAULT_API_URL, apiKey}
	ccaClient.Tasks = TaskApi{ccaClient.apiClient}
	return ccaClient
}

func NewCCAClientWithCustomURL(apiURL string, apiKey string) CCAClient {
	ccaClient := CCAClient{}
	ccaClient.apiClient = CCAApiClient{apiURL, apiKey}
	ccaClient.Tasks = TaskApi{ccaClient.apiClient}
	return ccaClient
}

//Get the Resources for a specific serviceCode and environmentName
//For now it assumes that the serviceCode belongs to a cloud.ca service
func (c CCAClient) GetResources(serviceCode string, environmentName string) {

}