package cloudca

import (
	"github.com/cloud-ca/go-cloudca/api"
)

type Resources struct {
	apiClient api.CCAApiClient
	serviceCode string
	environmentName string
	Instances InstanceService
}

func NewResources(apiClient api.CCAApiClient, serviceCode string, environmentName string) Resources {
	return Resources{
		apiClient: apiClient,
		serviceCode: serviceCode,
		environmentName: environmentName,
		Instances: NewInstanceService(apiClient, serviceCode, environmentName),
	}
}