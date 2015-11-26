package cloudca

import (
	"github.com/cloud-ca/go-cloudca/api"
)

type Resources struct {
	apiClient api.CcaApiClient
	serviceCode string
	environmentName string
	Instances InstanceService
	Templates TemplateService
	ComputeOfferings ComputeOfferingService
	DiskOfferings DiskOfferingService
	SSHKeys DiskOfferingService
}

func NewResources(apiClient api.CcaApiClient, serviceCode string, environmentName string) Resources {
	return Resources{
		apiClient: apiClient,
		serviceCode: serviceCode,
		environmentName: environmentName,
		Instances: NewInstanceService(apiClient, serviceCode, environmentName),
		Templates: NewTemplateService(apiClient, serviceCode, environmentName),
		ComputeOfferings: NewComputeOfferingService(apiClient, serviceCode, environmentName),
	}
}