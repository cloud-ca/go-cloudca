package cloudca

import (
	"github.com/cloud-ca/go-cloudca/services"
	"github.com/cloud-ca/go-cloudca/api"
	"encoding/json"
)

type Service struct {
	Name string `json:"name,omitempty"`
	Capabilities map[string]interface{} `json:"capabilities,omitempty"`
}

type Tier struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ZoneId string `json:"zoneid,omitempty"`
	ZoneName string `json:"zonename,omitempty"`
	Cidr string `json:"cidr,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	NetworkOfferingId string `json:"networkofferingid,omitempty"`
	IsSystem bool `json:"issystem,omitempty"`
	VpcId string `json:"vpcid,omitempty"`
	Domain string `json:"domain,omitempty"`
	DomainId string `json:"domainid,omitempty"`
	Project string `json:"project,omitempty"`
	ProjectId string `json:"projectid,omitempty"`
	AclId string `json:"aclId,omitempty"`
	Services []Service `json:"service,omitempty"`
}

type TierService interface {
	Get(id string) (*Tier, error)
	List() ([]Tier, error)
	ListOfVpc(vpcId string) ([]Tier, error)
	ListWithOptions(options map[string]string) ([]Tier, error)
}

type TierApi struct {
	entityService services.EntityService
}

func NewTierService(apiClient api.CcaClient, serviceCode string, environmentName string) TierService {
	return &TierApi{
		entityService: services.NewEntityService(apiClient, serviceCode, environmentName, TIER_ENTITY_TYPE),
	}
}

func parseTier(data []byte) *Tier {
	tier := Tier{}
	json.Unmarshal(data, &tier)
	return &tier
}

func parseTierList(data []byte) []Tier {
	tiers := []Tier{}
	json.Unmarshal(data, &tiers)
	return tiers
}

//Get tier with the specified id for the current environment
func (tierApi *TierApi) Get(id string) (*Tier, error) {
	data, err := tierApi.entityService.Get(id, map[string]string{})
	if err != nil {
		return nil, err
	}
	return parseTier(data), nil
}

//List all tiers for the current environment
func (tierApi *TierApi) List() ([]Tier, error) {
	return tierApi.ListWithOptions(map[string]string{})
}

//List all tiers of a vpc for the current environment
func (tierApi *TierApi) ListOfVpc(vpcId string) ([]Tier, error) {
	return tierApi.ListWithOptions(map[string]string{
			vpcId: vpcId,
		})
}

//List all tiers for the current environment. Can use options to do sorting and paging.
func (tierApi *TierApi) ListWithOptions(options map[string]string) ([]Tier, error) {
	data, err := tierApi.entityService.List(options)
	if err != nil {
		return nil, err
	}
	return parseTierList(data), nil
}
