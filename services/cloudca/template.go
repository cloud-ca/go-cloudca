package services/cloudca

type Template struct {

}

type TemplateService interface {
	Get(id string) (Template, error)
	GetByName(name string) (Template, error)
	List() ([]Template, error)
}

type TemplateApi struct {
	entityService EntityService
}

func NewInstanceService(apiClient CCAApiClient, serviceCode string, environmentName string) TemplateService {
	return TemplateApi{
		"entityService": NewEntityService(apiClient, serviceCode, environmentName, TEMPLATE_ENTITY_TYPE)
	}
}

func (templateApi TemplateApi) Get(id string) (Template, error) {
	return Template{}, nil
}

func (templateApi TemplateApi) GetByName(name string) (Template, error) {
	return Template{}, nil
}

func (templateApi TemplateApi) List() ([]Template, error) {
	return []Template{}, nil
}
