package cloudca

import (
	"github.com/cloud-ca/go-cloudca/services"
	"github.com/cloud-ca/go-cloudca/api"
	"encoding/json"
)

type Tier struct {

}

type TierService interface {
	Get(id string) (*Tier, error)
	List() ([]Tier, error)
	ListForVpc(vpcId string) ([]Tier, error)
	ListWithOptions(options map[string]string) ([]Tier, error)
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
	data, err := tierApi.entityService.Get(id, map[string]string{})
	if err != nil {
		return nil, err
	}
	tier := Tier{}
	json.Unmarshal(data, &tier)
	return &tier, nil
}

func (tierApi *TierApi) List() ([]Tier, error) {
	return tierApi.ListWithOptions(map[string]string{})
}


func (tierApi *TierApi) ListForVpc(vpcId string) ([]Tier, error) {
	return tierApi.ListWithOptions(map[string]string{
			vpcId: vpcId,
		})
}

func (tierApi *TierApi) ListWithOptions(options map[string]string) ([]Tier, error) {
	data, err := tierApi.entityService.List(options)
	if err != nil {
		return nil, err
	}
	tiers := []Tier{}
	json.Unmarshal(data, &tiers)
	return tiers, nil
}
