package services/cloudca

type ComputeOffering struct {

}

type ComputeOfferingService interface {
	Get(id string) (ComputeOffering, error)
	GetByName(name string) (ComputeOffering, error)
	List() ([]ComputeOffering, error)
}

type ComputeOfferingApi struct {
	entityService EntityService
}

func NewInstanceService(apiClient CCAApiClient, serviceCode string, environmentName string) ComputeOfferingService {
	return ComputeOfferingApi{
		"entityService": NewEntityService(apiClient, serviceCode, environmentName, COMPUTE_OFFERING_ENTITY_TYPE)
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
