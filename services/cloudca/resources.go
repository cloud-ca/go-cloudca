package services/cloudca

type Resources struct {
	apiClient CCAApiClient
	serviceCode string
	environmentName string
	Instances InstanceService
}