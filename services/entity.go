package services

import (
	"github.com/cloud-ca/go-cloudca/api"
)

type EntityService interface {
	Get(id string, options map[string]string) ([]byte, error)
	List(options map[string]string) ([]byte, error)
	Execute(operation string, body []byte, options map[string]string) ([]byte, error)
	Create(body []byte, options map[string]string) ([]byte, error)
	Update(body []byte, options map[string]string) ([]byte, error)
	Delete(body []byte, options map[string]string) ([]byte, error)
}

type EntityApi struct {
	apiClient api.CCAApiClient
	taskService TaskService
	serviceCode string
	environmentName string
	entityType string
}

func NewEntityService(apiClient api.CCAApiClient, serviceCode string, environmentName string, entityType string) EntityService {
	return EntityApi{
		apiClient: apiClient,
		taskService: NewTaskService(apiClient),
		serviceCode: serviceCode,
		environmentName: environmentName,
		entityType: entityType,
	}
}

func (entityApi EntityApi) buildEndpoint() string {
	return "/" + entityApi.serviceCode + "/" + entityApi.environmentName + "/" + entityApi.entityType
}

func (entityApi EntityApi) Get(id string, options map[string]string) ([]byte, error) {
	return nil, nil
}

func (entityApi EntityApi) List(options map[string]string) ([]byte, error) {
	return nil, nil
}

func (entityApi EntityApi) Execute(operation string, body []byte, options map[string]string) ([]byte, error) {
	optionsCopy := map[string]string{}
	for k, v := range options {
		optionsCopy[k] = v
	}
	optionsCopy["operation"] = operation
	return nil, nil
}

func (entityApi EntityApi) Create(body []byte, options map[string]string) ([]byte, error) {
	return nil, nil
}

func (entityApi EntityApi) Update(body []byte, options map[string]string) ([]byte, error) {
	return nil, nil
}

func (entityApi EntityApi) Delete(body []byte, options map[string]string) ([]byte, error) {
	return nil, nil
}
