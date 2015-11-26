package services/cloudca

type SSHKey struct {

}

type SSHKeyService interface {
	Get(name string) (SSHKey, error)
	List() ([]SSHKey, error)
}

type SSHKeyApi struct {
	entityService EntityService
}

func NewInstanceService(apiClient CCAApiClient, serviceCode string, environmentName string) SSHKeyService {
	return SSHKeyApi{
		"entityService": NewEntityService(apiClient, serviceCode, environmentName, SSH_KEY_ENTITY_TYPE)
	}
}

func (sshKeyApi SSHKeyApi) Get(name string) (SSHKey, error) {
	return SSHKey{}, nil
}

func (sshKeyApi SSHKeyApi) List() ([]SSHKey, error) {
	return []SSHKey{}, nil
}
