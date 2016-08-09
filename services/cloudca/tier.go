package cloudca

import (
	"encoding/json"
	"github.com/cloud-ca/go-cloudca/api"
	"github.com/cloud-ca/go-cloudca/services"
)

type Service struct {
	Name         string                 `json:"name,omitempty"`
	Capabilities map[string]interface{} `json:"capabilities,omitempty"`
}

type Tier struct {
	Id                string    `json:"id,omitempty"`
	Name              string    `json:"name,omitempty"`
	Description       string    `json:"description,omitempty"`
	VpcId             string    `json:"vpcId,omitempty"`
	NetworkOfferingId string    `json:"networkOfferingId,omitempty"`
	NetworkAclId      string    `json:"networkAclId,omitempty"`
	ZoneId            string    `json:"zoneid,omitempty"`
	ZoneName          string    `json:"zonename,omitempty"`
	Cidr              string    `json:"cidr,omitempty"`
	Type              string    `json:"type,omitempty"`
	State             string    `json:"state,omitempty"`
	Gateway           string    `json:"gateway,omitempty"`
	IsSystem          bool      `json:"issystem,omitempty"`
	Domain            string    `json:"domain,omitempty"`
	DomainId          string    `json:"domainid,omitempty"`
	Project           string    `json:"project,omitempty"`
	ProjectId         string    `json:"projectid,omitempty"`
	Services          []Service `json:"service,omitempty"`
}

type TierService interface {
	Get(id string) (*Tier, error)
	List() ([]Tier, error)
	ListOfVpc(vpcId string) ([]Tier, error)
	ListWithOptions(options map[string]string) ([]Tier, error)
	Create(tier Tier, options map[string]string) (*Tier, error)
	Update(id string, tier Tier) (*Tier, error)
	Delete(id string) (bool, error)
	ChangeAcl(id string, aclId string) (bool, error)
}

type TierApi struct {
	entityService services.EntityService
}

func NewTierService(apiClient api.ApiClient, serviceCode string, environmentName string) TierService {
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

func (tierApi *TierApi) Create(tier Tier, options map[string]string) (*Tier, error) {
	send, merr := json.Marshal(tier)
	if merr != nil {
		return nil, merr
	}
	body, err := tierApi.entityService.Create(send, options)
	if err != nil {
		return nil, err
	}
	return parseTier(body), nil
}

func (tierApi *TierApi) Update(id string, tier Tier) (*Tier, error) {
	send, merr := json.Marshal(tier)
	if merr != nil {
		return nil, merr
	}
	body, err := tierApi.entityService.Update(id, send, map[string]string{})
	if err != nil {
		return nil, err
	}
	return parseTier(body), nil
}

func (tierApi *TierApi) Delete(id string) (bool, error) {
	_, err := tierApi.entityService.Delete(id, []byte{}, map[string]string{})
	return err == nil, err
}

func (tierApi *TierApi) ChangeAcl(id string, aclId string) (bool, error) {
	send, merr := json.Marshal(Tier{
		NetworkAclId: aclId,
	})
	if merr != nil {
		return false, merr
	}
	_, err := tierApi.entityService.Execute(id, "replace", send, map[string]string{})
	return err == nil, err
}
