package cloudca

import (
	"github.com/cloud-ca/go-cloudca/services"
	"github.com/cloud-ca/go-cloudca/api"
	"encoding/json"
)

type ComputeOffering struct {

}

type ComputeOfferingService interface {
	Get(id string) (*ComputeOffering, error)
	List() ([]ComputeOffering, error)
	ListWithOptions(options map[string]string) ([]ComputeOffering, error)
}

type ComputeOfferingApi struct {
	entityService services.EntityService
}

func NewComputeOfferingService(apiClient api.CcaApiClient, serviceCode string, environmentName string) ComputeOfferingService {
	return &ComputeOfferingApi{
		entityService: services.NewEntityService(apiClient, serviceCode, environmentName, COMPUTE_OFFERING_ENTITY_TYPE),
	}
}

func (computeOfferingApi *ComputeOfferingApi) Get(id string) (*ComputeOffering, error) {
	data, err := computeOfferingApi.entityService.Get(id, map[string]string{})
	if err != nil {
		return nil, err
	}
	computeOffering := ComputeOffering{}
	json.Unmarshal(data, &computeOffering)
	return &computeOffering, nil
}

func (computeOfferingApi *ComputeOfferingApi) List() ([]ComputeOffering, error) {
	return computeOfferingApi.ListWithOptions(map[string]string{})
}

func (computeOfferingApi *ComputeOfferingApi) ListWithOptions(options map[string]string) ([]ComputeOffering, error) {
	data, err := computeOfferingApi.entityService.List(options)
	if err != nil {
		return nil, err
	}
	computeOfferings := []ComputeOffering{}
	json.Unmarshal(data, &computeOfferings)
	return computeOfferings, nil
}
