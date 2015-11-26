package services/cloudca

type Tier struct {

}

type TierService interface {
	Get(id string) (Tier, error)
	GetByName(name string) (Tier, error)
	List() ([]Tier, error)
}

type TierApi struct {
	entityService EntityService
}

func NewInstanceService(apiClient CCAApiClient, serviceCode string, environmentName string) TierService {
	return TierApi{
		"entityService": NewEntityService(apiClient, serviceCode, environmentName, TIER_ENTITY_TYPE)
	}
}

func (tierApi TierApi) Get(id string) (Tier, error) {
	return Tier{}, nil
}

func (tierApi TierApi) GetByName(name string) (Tier, error) {
	return Tier{}, nil
}

func (tierApi TierApi) List() ([]Tier, error) {
	return []Tier{}, nil
}
