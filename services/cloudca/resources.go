package cloudca

import (
	"github.com/cloud-ca/go-cloudca/api"
)

const (
	CLOUD_CA_SERVICE = "cloudca"
)

type Resources struct {
	apiClient        api.ApiClient
	serviceCode      string
	environmentName  string
	Instances        InstanceService
	Volumes          VolumeService
	Templates        TemplateService
	ComputeOfferings ComputeOfferingService
	DiskOfferings    DiskOfferingService
	SSHKeys          SSHKeyService
	Tiers            TierService
	Vpcs             VpcService
	VpcOfferings     VpcOfferingService
	NetworkOfferings NetworkOfferingService
	PublicIps        PublicIpService
	NetworkAcls      NetworkAclService
	NetworkAclRules  NetworkAclRuleService
	Zones            ZoneService
	PortForwardingRules PortForwardingRuleService
}

func NewResources(apiClient api.ApiClient, serviceCode string, environmentName string) Resources {
	return Resources{
		apiClient:           apiClient,
		serviceCode:         serviceCode,
		environmentName:     environmentName,
		Instances:           NewInstanceService(apiClient, serviceCode, environmentName),
		Volumes:             NewVolumeService(apiClient, serviceCode, environmentName),
		Templates:           NewTemplateService(apiClient, serviceCode, environmentName),
		ComputeOfferings:    NewComputeOfferingService(apiClient, serviceCode, environmentName),
		DiskOfferings:       NewDiskOfferingService(apiClient, serviceCode, environmentName),
		Tiers:               NewTierService(apiClient, serviceCode, environmentName),
		Vpcs:                NewVpcService(apiClient, serviceCode, environmentName),
		VpcOfferings:        NewVpcOfferingService(apiClient, serviceCode, environmentName),
		NetworkOfferings:    NewNetworkOfferingService(apiClient, serviceCode, environmentName),
		NetworkAcls:         NewNetworkAclService(apiClient, serviceCode, environmentName),
		NetworkAclRules:     NewNetworkAclRuleService(apiClient, serviceCode, environmentName),
		PublicIps:           NewPublicIpService(apiClient, serviceCode, environmentName),
		PortForwardingRules: NewPortForwardingRuleService(apiClient, serviceCode, environmentName),
		Zones:               NewZoneService(apiClient, serviceCode, environmentName),
	}
}

func (resources Resources) GetServiceType() string {
	return CLOUD_CA_SERVICE
}
