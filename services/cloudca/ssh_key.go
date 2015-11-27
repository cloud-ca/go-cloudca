package cloudca

import (
	"github.com/cloud-ca/go-cloudca/services"
	"github.com/cloud-ca/go-cloudca/api"
	"encoding/json"
)

type SSHKey struct {
	Name string `json:"name,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
}

type SSHKeyService interface {
	Get(name string) (*SSHKey, error)
	List() ([]SSHKey, error)
	ListWithOptions(options map[string]string) ([]SSHKey, error)
}

type SSHKeyApi struct {
	entityService services.EntityService
}

func NewSSHKeyService(apiClient api.CcaApiClient, serviceCode string, environmentName string) SSHKeyService {
	return &SSHKeyApi{
		entityService: services.NewEntityService(apiClient, serviceCode, environmentName, SSH_KEY_ENTITY_TYPE),
	}
}

func (sshKeyApi *SSHKeyApi) Get(name string) (*SSHKey, error) {
	return nil, nil
}

func (sshKeyApi *SSHKeyApi) List() ([]SSHKey, error) {
	return sshKeyApi.ListWithOptions(map[string]string{})
}

func (sshKeyApi *SSHKeyApi) ListWithOptions(options map[string]string) ([]SSHKey, error) {
	data, err := sshKeyApi.entityService.List(options)
	if err != nil {
		return nil, err
	}
	sshKeys := []SSHKey{}
	json.Unmarshal(data, &sshKeys)
	return sshKeys, nil
}
