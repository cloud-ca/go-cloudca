package main

type EntityApi struct {
	apiURL string
	apiKey string
	serviceCode string
	environmentName string
}



func (e EntityApi) Execute(entityType string, operation string, parameters map[string]string, body string) {

}

func (e EntityApi) Create(entityType string, parameters map[string]string, body string) {

}

func (e EntityApi) Update(entityType string, parameters map[string]string, body string) {

}

func (e EntityApi) Delete(entityType string, parameters map[string]string, body string) {

}