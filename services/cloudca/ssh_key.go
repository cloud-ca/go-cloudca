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

func NewSSHKeyService(apiClient api.ApiClient, serviceCode string, environmentName string) SSHKeyService {
	return &SSHKeyApi{
		entityService: services.NewEntityService(apiClient, serviceCode, environmentName, SSH_KEY_ENTITY_TYPE),
	}
}

func parseSSHKey(data []byte) *SSHKey {
	sshKey := SSHKey{}
	json.Unmarshal(data, &sshKey)
	return &sshKey
}

func parseSSHKeyList(data []byte) []SSHKey {
	sshKeys := []SSHKey{}
	json.Unmarshal(data, &sshKeys)
	return sshKeys
}

//Get SSH key with the specified id for the current environment
func (sshKeyApi *SSHKeyApi) Get(name string) (*SSHKey, error) {
	data, err := sshKeyApi.entityService.Get(name, map[string]string{})
	if err != nil {
		return nil, err
	}
	return parseSSHKey(data), nil
}

//List all SSH keys for the current environment
func (sshKeyApi *SSHKeyApi) List() ([]SSHKey, error) {
	return sshKeyApi.ListWithOptions(map[string]string{})
}

//List all SSH keys for the current environment. Can use options to do sorting and paging.
func (sshKeyApi *SSHKeyApi) ListWithOptions(options map[string]string) ([]SSHKey, error) {
	data, err := sshKeyApi.entityService.List(options)
	if err != nil {
		return nil, err
	}
	return parseSSHKeyList(data), nil
}
