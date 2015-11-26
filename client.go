package main

type CCAClient struct {
	apiClient CCAApiClient
	Tasks TaskService
}

func NewCCAClient(apiURL string, apiKey string) CCAClient {
	ccaClient := CCAClient{}
	ccaClient.apiClient = CCAApiClient{apiURL, apiKey}
	ccaClient.Tasks = TaskApi{ccaClient.apiClient}
	return ccaClient
}

//Get the Resources for a specific serviceCode and environmentName
func (c CCAClient) GetResources(serviceCode string, environmentName string) {

}