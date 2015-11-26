package services/cloudca

type Instance struct {
	Id string
}

type InstanceService interface {
}

type InstanceApi struct {
	entityService EntityService
}

func NewInstanceService(apiClient CCAApiClient, serviceCode string, environmentName string) *InstanceService {
	return &InstanceApi{
		"entityService": NewEntityService(apiClient, serviceCode, environmentName, INSTANCE_ENTITY_TYPE)
	}
}