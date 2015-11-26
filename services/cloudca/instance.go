package services/cloudca

type Instance struct {
	Id string
}

type InstanceService interface {
	Get(id string) (Instance, error)
	GetByName(name string) (Instance, error)
	List() ([]Instance, error)
	Create(Instance)
}

type InstanceApi struct {
	entityService EntityService
}

func NewInstanceService(apiClient CCAApiClient, serviceCode string, environmentName string) *InstanceService {
	return &InstanceApi{
		"entityService": NewEntityService(apiClient, serviceCode, environmentName, INSTANCE_ENTITY_TYPE)
	}
}

func (instanceApi InstanceApi) Get(id string) (Instance, error) {
	return Instance{}, nil
}

func (instanceApi InstanceApi) GetByName(name string) (Instance, error) {
	return Instance{}, nil
}


func (instanceApi InstanceApi) GetByName(name string) (Instance, error) {
	return Instance{}, nil
}


func (instanceApi InstanceApi) List() ([]Instance, error) {
	return []Instance{}, nil
}
