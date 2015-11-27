package cloudca

import (
	"github.com/cloud-ca/go-cloudca/services"
	"github.com/cloud-ca/go-cloudca/api"
	"encoding/json"
)

type DiskOffering struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	GbSize int `json:"gbSize,omitempty"`
	StorageTier string `json:"storageTier,omitempty"`
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

//Get disk offering with the specified id for the current environment
func (diskOfferingApi *DiskOfferingApi) Get(id string) (*DiskOffering, error) {
	data, err := diskOfferingApi.entityService.Get(id, map[string]string{})
	if err != nil {
		return nil, err
	}
	diskOffering := DiskOffering{}
	json.Unmarshal(data, &diskOffering)
	return &diskOffering, nil
}

//List all disk offerings for the current environment
func (diskOfferingApi *DiskOfferingApi) List() ([]DiskOffering, error) {
	return diskOfferingApi.ListWithOptions(map[string]string{})
}

//List all disk offerings for the current environment. Can use options to do sorting and paging.
func (diskOfferingApi *DiskOfferingApi) ListWithOptions(options map[string]string) ([]DiskOffering, error) {
	data, err := diskOfferingApi.entityService.List(options)
	if err != nil {
		return nil, err
	}
	diskOfferings := []DiskOffering{}
	json.Unmarshal(data, &diskOfferings)
	return diskOfferings, nil
}
