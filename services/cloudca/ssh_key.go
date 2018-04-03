package cloudca

import (
  "encoding/json"
  "github.com/cloud-ca/go-cloudca/api"
  "github.com/cloud-ca/go-cloudca/services"
)

type SSHKey struct {
  Name        string `json:"name,omitempty"`
  PublicKey   string `json:"publicKey,omitempty"`
  Fingerprint string `json:"fingerprint,omitempty"`
}

type SSHKeyService interface {
  Get(name string) (*SSHKey, error)
  List() ([]SSHKey, error)
  ListWithOptions(options map[string]string) ([]SSHKey, error)
  Register(sshKey SSHKey) (*SSHKey, error)
  Delete(name string) (bool, error)
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

//Get SSH key with the specified name for the current environment
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

//Register a ssh key for the current environment
func (sshKeyApi *SSHKeyApi) Register(sshKey SSHKey) (*SSHKey, error) {
  msg, err := json.Marshal(sshKey)
  if err != nil {
    return nil, err
  }
  result, err := sshKeyApi.entityService.Create(msg, map[string]string{})
  if err != nil {
  return nil, err
  }
  return parseSSHKey(result), nil
}

//Delete a ssh key in the current environment
func (sshKeyApi *SSHKeyApi) Delete(name string) (bool, error) {
  _, err := sshKeyApi.entityService.Delete(name, []byte{}, map[string]string{})
  return err == nil, err
}
