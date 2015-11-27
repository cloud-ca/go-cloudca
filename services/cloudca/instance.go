package cloudca

import (
	"github.com/cloud-ca/go-cloudca/services"
	"github.com/cloud-ca/go-cloudca/api"
	"encoding/json"
)

type Instance struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	State string `json:"state,omitempty"`
	TemplateId string `json:"templateId,omitempty"`
	TemplateName string `json:"templateName,omitempty"`
	IsPasswordEnabled bool `json:"isPasswordEnabled,omitempty"`
	IsSshKeyEnabled bool `json:"isSshKeyEnabled,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	SSHKeyName string `json:"sshKeyName,omitempty"`
	ComputeOfferingId string `json:"computeOfferingId,omitempty"`
	ComputeOfferingName string `json:"computeOfferingName,omitempty"`
	CpuCount int `json:"cpuCount,omitempty"`
	MemoryInMB int `json:"memoryInMB,omitempty"`
	ZoneId string `json:"zoneId,omitempty"`
	ZoneName string `json:"zoneName,omitempty"`
	ProjectId string `json:"projectId,omitempty"`
	NetworkId string `json:"networkId,omitempty"`
	NetworkName string `json:"networkName,omitempty"`
	MacAddress string `json:"macAddress,omitempty"`
	UserData string `json:"userData,omitempty"`
	PublicIps []PublicIp `json:"publicIPs,omitempty"`
}

type InstanceService interface {
	Get(id string) (*Instance, error)
	List() ([]Instance, error)
	ListWithOptions(options map[string]string) ([]Instance, error)
	Create(Instance) (*Instance, error)
	Delete(id string, purge bool) (bool, error)
	Exists(id string) (bool, error)
}

type InstanceApi struct {
	entityService services.EntityService
}

func NewInstanceService(apiClient api.CcaApiClient, serviceCode string, environmentName string) InstanceService {
	return &InstanceApi{
		entityService: services.NewEntityService(apiClient, serviceCode, environmentName, INSTANCE_ENTITY_TYPE),
	}
}

func (instanceApi *InstanceApi) Get(id string) (*Instance, error) {
	data, err := instanceApi.entityService.Get(id, map[string]string{})
	if err != nil {
		return nil, err
	}
	instance := Instance{}
	json.Unmarshal(data, &instance)
	return &instance, nil
}

func (instanceApi *InstanceApi) List() ([]Instance, error) {
	return instanceApi.ListWithOptions(map[string]string{})
}

func (instanceApi *InstanceApi) ListWithOptions(options map[string]string) ([]Instance, error) {
	data, err := instanceApi.entityService.List(options)
	if err != nil {
		return nil, err
	}
	instances := []Instance{}
	json.Unmarshal(data, &instances)
	return instances, nil
}

func (instanceApi *InstanceApi) Create(instance Instance) (*Instance, error) {
	return nil, nil
}

func (instanceApi *InstanceApi) Delete(id string, purge bool) (bool, error) {
	return false, nil
}

func (instanceApi *InstanceApi) Exists(id string) (bool, error) {
	return false, nil
}
