package cloudca

import (
	"github.com/cloud-ca/go-cloudca/services"
	"github.com/cloud-ca/go-cloudca/api"
	"encoding/json"
)

type Template struct {

}

type TemplateService interface {
	Get(id string) (*Template, error)
	List() ([]Template, error)
	ListWithOptions(options map[string]string) ([]Template, error)
}

type TemplateApi struct {
	entityService services.EntityService
}

func NewTemplateService(apiClient api.CcaApiClient, serviceCode string, environmentName string) TemplateService {
	return &TemplateApi{
		entityService: services.NewEntityService(apiClient, serviceCode, environmentName, TEMPLATE_ENTITY_TYPE),
	}
}

func (templateApi *TemplateApi) Get(id string) (*Template, error) {
	data, err := templateApi.entityService.Get(id, map[string]string{})
	if err != nil {
		return nil, err
	}
	template := Template{}
	json.Unmarshal(data, &template)
	return &template, nil
}

func (templateApi *TemplateApi) List() ([]Template, error) {
	return templateApi.ListWithOptions(map[string]string{})
}

func (templateApi *TemplateApi) ListWithOptions(options map[string]string) ([]Template, error) {
	data, err := templateApi.entityService.List(options)
	if err != nil {
		return nil, err
	}
	templates := []Template{}
	json.Unmarshal(data, &templates)
	return templates, nil
}