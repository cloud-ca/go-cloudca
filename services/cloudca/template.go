package cloudca

import (
	"github.com/cloud-ca/go-cloudca/services"
	"github.com/cloud-ca/go-cloudca/api"
	"encoding/json"
)

type Template struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Size int `json:"size,omitempty"`
	IsPublic bool `json:"isPublic,omitempty"`
	IsReady bool `json:"isReady,omitempty"`
	SSHKeyEnabled bool `json:"sshKeyEnabled,omitempty"`
	Extractable bool `json:"extractable,omitempty"`
	OSType string `json:"osType,omitempty"`
	OSTypeId string `json:"osTypeId,omitempty"`
	Hypervisor string `json:"hypervisor,omitempty"`
	Format string `json:"format,omitempty"`
	ZoneName string `json:"zoneName,omitempty"`
	ProjectId string `json:"projectId,omitempty"`
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

//Get template with the specified id for the current environment
func (templateApi *TemplateApi) Get(id string) (*Template, error) {
	data, err := templateApi.entityService.Get(id, map[string]string{})
	if err != nil {
		return nil, err
	}
	template := Template{}
	json.Unmarshal(data, &template)
	return &template, nil
}

//List all templates for the current environment
func (templateApi *TemplateApi) List() ([]Template, error) {
	return templateApi.ListWithOptions(map[string]string{})
}

//List all templates for the current environment. Can use options to do sorting and paging.
func (templateApi *TemplateApi) ListWithOptions(options map[string]string) ([]Template, error) {
	data, err := templateApi.entityService.List(options)
	if err != nil {
		return nil, err
	}
	templates := []Template{}
	json.Unmarshal(data, &templates)
	return templates, nil
}