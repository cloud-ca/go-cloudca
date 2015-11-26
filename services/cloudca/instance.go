package services/cloudca

type Instance struct {
	Id string
}

type InstanceService interface {
}

type InstanceApi struct {
	apiClient CCAApiClient
	serviceCode string
	environmentName string
}