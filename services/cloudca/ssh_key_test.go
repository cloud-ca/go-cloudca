package cloudca

import (
	"testing"
	"github.com/cloud-ca/go-cloudca/mocks"
	"github.com/cloud-ca/go-cloudca/mocks/services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	SSH_KEY_NAME = "test_ssh_key"
	SSH_KEY_FINGERPRINT = "test_fingerprint"
)

func buildGetSSHKeySuccessResponse() []byte {
	return  []byte(`{"name": "` + SSH_KEY_NAME + 
			`","fingerprint":"` + SSH_KEY_FINGERPRINT + `"}`)
}

func TestGetSSHKeyReturnSSHKeyIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	sshKeyService := SSHKeyApi{
		entityService: mockEntityService,
	}

	expectedSSHKey := SSHKey{Name: SSH_KEY_NAME, 
							 Fingerprint: SSH_KEY_FINGERPRINT}

	mockEntityService.EXPECT().Get(SSH_KEY_NAME, gomock.Any()).Return(buildGetSSHKeySuccessResponse(), nil)

	//when
	sshKey, _ := sshKeyService.Get(SSH_KEY_NAME)

	//then
	if assert.NotNil(t, sshKey) {
		assert.Equal(t, expectedSSHKey, *sshKey)
	}
}

func TestGetSSHKeyReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	sshKeyService := SSHKeyApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(SSH_KEY_NAME, gomock.Any()).Return(nil, mockError)

	//when
	sshKey, err := sshKeyService.Get(SSH_KEY_NAME)

	//then
	assert.Nil(t, sshKey)
	assert.Equal(t, mockError, err)

}

func buildListSSHKeySuccessResponse() []byte {
	return []byte(`[
		{"name": "list_name_1", "fingerprint": "list_fingerprint_1"},
		{"name": "list_name_2", "fingerprint": "list_fingerprint_2"}
	]`)
}

func TestListSSHKeyReturnSSHKeysIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	sshKeyService := SSHKeyApi{
		entityService: mockEntityService,
	}

	expectedSSHKeys := []SSHKey{
		SSHKey{
			Name: "list_name_1",
			Fingerprint: "list_fingerprint_1",
		},
		SSHKey{
			Name: "list_name_2",
			Fingerprint: "list_fingerprint_2",
		},
	}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListSSHKeySuccessResponse(), nil)

	//when
	sshKeys, _ := sshKeyService.List()

	//then
	if assert.NotNil(t, sshKeys) {
		assert.Equal(t, expectedSSHKeys, sshKeys)
	}
}


func TestListSSHKeyReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	sshKeyService := SSHKeyApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	sshKeys, err := sshKeyService.List()

	//then
	assert.Nil(t, sshKeys)
	assert.Equal(t, mockError, err)

}