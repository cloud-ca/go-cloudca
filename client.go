package gocca


const (
	DEFAULT_API_URL = "https://api.cloud.ca/v1/"
)

type CCAClient struct {
	apiURL string
	apiKey string
	apiClient CCAApiClient
	Tasks TaskService
}

func NewCCAClient(apiKey string) CCAClient {
	return NewCCAClientWithCustomURL(DEFAULT_API_URL, apiKey)
}

func NewCCAClientWithCustomURL(apiURL string, apiKey string) CCAClient {
	ccaClient := CCAClient{}
	ccaClient.apiURL = apiURL
	ccaClient.apiKey = apiKey
	ccaClient.apiClient = CCAApiClient{apiURL, apiKey}
	ccaClient.Tasks = TaskApi{ccaClient.apiClient}
	return ccaClient
}

//Get the Resources for a specific serviceCode and environmentName
//For now it assumes that the serviceCode belongs to a cloud.ca service
func (c CCAClient) GetResources(serviceCode string, environmentName string) {

}

func (c CCAClient) GetApiURL() string {
	return c.apiURL
}

func (c CCAClient) GetApiKey() string {
	return c.apiKey
}