package cloudca

import (
	"github.com/cloud-ca/go-cloudca/services"
	"github.com/cloud-ca/go-cloudca/api"
)

type SSHKey struct {

}

type SSHKeyService interface {
	Get(name string) (SSHKey, error)
	List() ([]SSHKey, error)
}

type SSHKeyApi struct {
	entityService services.EntityService
}

func NewSSHKeyService(apiClient api.CCAApiClient, serviceCode string, environmentName string) SSHKeyService {
	return SSHKeyApi{
		entityService: services.NewEntityService(apiClient, serviceCode, environmentName, SSH_KEY_ENTITY_TYPE),
	}
}

func (sshKeyApi SSHKeyApi) Get(name string) (SSHKey, error) {
	return SSHKey{}, nil
}

func (sshKeyApi SSHKeyApi) List() ([]SSHKey, error) {
	return []SSHKey{}, nil
}
