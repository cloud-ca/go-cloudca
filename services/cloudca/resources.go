package cloudca

import (
	"github.com/cloud-ca/go-cloudca/api"
)

const (
	CLOUD_CA_SERVICE = "cloudca"
)

type Resources struct {
	apiClient api.CcaApiClient
	serviceCode string
	environmentName string
	Instances InstanceService
	Templates TemplateService
	ComputeOfferings ComputeOfferingService
	DiskOfferings DiskOfferingService
	SSHKeys SSHKeyService
	Tiers TierService
}

func NewResources(apiClient api.CcaApiClient, serviceCode string, environmentName string) Resources {
	return Resources{
		apiClient: apiClient,
		serviceCode: serviceCode,
		environmentName: environmentName,
		Instances: NewInstanceService(apiClient, serviceCode, environmentName),
		Templates: NewTemplateService(apiClient, serviceCode, environmentName),
		ComputeOfferings: NewComputeOfferingService(apiClient, serviceCode, environmentName),
		Tiers: NewTierService(apiClient, serviceCode, environmentName),
		
	}
}

func (resources Resources) GetServiceType() string {
	return CLOUD_CA_SERVICE
}