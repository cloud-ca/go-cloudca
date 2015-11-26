package cloudca

import (
	"github.com/cloud-ca/go-cloudca/services"
	"github.com/cloud-ca/go-cloudca/api"
)

type ComputeOffering struct {

}

type ComputeOfferingService interface {
	Get(id string) (ComputeOffering, error)
	GetByName(name string) (ComputeOffering, error)
	List() ([]ComputeOffering, error)
}

type ComputeOfferingApi struct {
	entityService services.EntityService
}

func NewComputeOfferingService(apiClient api.CcaApiClient, serviceCode string, environmentName string) ComputeOfferingService {
	return ComputeOfferingApi{
		entityService: services.NewEntityService(apiClient, serviceCode, environmentName, COMPUTE_OFFERING_ENTITY_TYPE),
	}
}

func (computeOfferingAPi ComputeOfferingApi) Get(id string) (ComputeOffering, error) {
	return ComputeOffering{}, nil
}

func (computeOfferingAPi ComputeOfferingApi) GetByName(name string) (ComputeOffering, error) {
	return ComputeOffering{}, nil
}

func (computeOfferingAPi ComputeOfferingApi) List() ([]ComputeOffering, error) {
	return []ComputeOffering{}, nil
}
