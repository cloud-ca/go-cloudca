package cloudca

import (
	"github.com/cloud-ca/go-cloudca/services"
	"github.com/cloud-ca/go-cloudca/api"
	"encoding/json"
)

type DiskOffering struct {

}

type DiskOfferingService interface {
	Get(id string) (*DiskOffering, error)
	List() ([]DiskOffering, error)
	ListWithOptions(options map[string]string) ([]DiskOffering, error)
}

type DiskOfferingApi struct {
	entityService services.EntityService
}

func NewDiskOfferingService(apiClient api.CcaApiClient, serviceCode string, environmentName string) DiskOfferingService {
	return &DiskOfferingApi{
		entityService: services.NewEntityService(apiClient, serviceCode, environmentName, DISK_OFFERING_ENTITY_TYPE),
	}
}

func (diskOfferingApi *DiskOfferingApi) Get(id string) (*DiskOffering, error) {
	data, err := diskOfferingApi.entityService.Get(id, map[string]string{})
	if err != nil {
		return nil, err
	}
	diskOffering := DiskOffering{}
	json.Unmarshal(data, &diskOffering)
	return &diskOffering, nil
}

func (diskOfferingApi *DiskOfferingApi) List() ([]DiskOffering, error) {
	return diskOfferingApi.ListWithOptions(map[string]string{})
}

func (diskOfferingApi *DiskOfferingApi) ListWithOptions(options map[string]string) ([]DiskOffering, error) {
	data, err := diskOfferingApi.entityService.List(options)
	if err != nil {
		return nil, err
	}
	diskOfferings := []DiskOffering{}
	json.Unmarshal(data, &diskOfferings)
	return diskOfferings, nil
}
