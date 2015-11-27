package cloudca

import (
	"github.com/cloud-ca/go-cloudca/services"
	"github.com/cloud-ca/go-cloudca/api"
)

type Tier struct {

}

type TierService interface {
	Get(id string) (*Tier, error)
	GetByName(name string) (*Tier, error)
	List() ([]Tier, error)
	ListForVpc(vpcId string) ([]Tier, error)
}

type TierApi struct {
	entityService services.EntityService
}

func NewTierService(apiClient api.CcaApiClient, serviceCode string, environmentName string) TierService {
	return &TierApi{
		entityService: services.NewEntityService(apiClient, serviceCode, environmentName, TIER_ENTITY_TYPE),
	}
}

func (tierApi *TierApi) Get(id string) (*Tier, error) {
	return nil, nil
}

func (tierApi *TierApi) GetByName(name string) (*Tier, error) {
	return nil, nil
}

func (tierApi TierApi) List() ([]Tier, error) {
	return nil, nil
}


func (tierApi TierApi) ListForVpc(vpcId string) ([]Tier, error) {
	return nil, nil
}

