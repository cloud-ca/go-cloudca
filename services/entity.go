package main

type EntityService interface {
	Find(id string, options map[string]string) ([]byte, error)
	List(options map[string]string) ([]byte, error)
	Execute(operation string, body []byte, options map[string]string) ([]byte, error)
	Create(body []byte, options map[string]string) ([]byte, error)
	Update(body []byte, options map[string]string) ([]byte, error)
	Delete(body []byte, options map[string]string) ([]byte, error)
}

type EntityApi struct {
	apiClient CCAApiClient
	serviceCode string
	environmentName string
	entityType string
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
