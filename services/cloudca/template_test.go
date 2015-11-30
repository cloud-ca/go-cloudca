package cloudca

import (
	"testing"
	"github.com/cloud-ca/go-cloudca/mocks"
	"github.com/cloud-ca/go-cloudca/mocks/services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"strconv"
)

const (
	TEMPLATE_ID = "test_template_id"
	TEMPLATE_NAME = "test_template"
	TEMPLATE_DESCRIPTION = "test_template_description"
	TEMPLATE_SIZE = 60
	TEMPLATE_IS_PUBLIC = true
	TEMPLATE_IS_READY = true
	TEMPLATE_SSH_KEY_ENABLED = false
	TEMPLATE_EXTRACTABLE = true
	TEMPLATE_OS_TYPE = "test_template_os_type"
	TEMPLATE_OS_TYPE_ID = "test_template_os_type_id"
	TEMPLATE_HYPERVISOR = "test_template_hypervisor"
	TEMPLATE_FORMAT = "test_template_format"
	TEMPLATE_ZONE_NAME = "test_template_zone_name"
	TEMPLATE_PROJECT_ID = "test_template_project_id"
)

func buildGetTemplateSuccessResponse() []byte {
	return  []byte(`{"id":"` + TEMPLATE_ID + `",` +
				   ` "name": "` + TEMPLATE_NAME + `",` +
				   ` "description": "` + TEMPLATE_DESCRIPTION + `",` +
				   ` "size": ` + strconv.Itoa(TEMPLATE_SIZE) + `,` +
				   ` "isPublic": ` + strconv.FormatBool(TEMPLATE_IS_PUBLIC) + `,` +
				   ` "isReady": ` + strconv.FormatBool(TEMPLATE_IS_READY) + `,` +
				   ` "sshKeyEnabled": ` + strconv.FormatBool(TEMPLATE_SSH_KEY_ENABLED) + `,` +
				   ` "extractable": ` + strconv.FormatBool(TEMPLATE_EXTRACTABLE) + `,` +
				   ` "osType": "` + TEMPLATE_OS_TYPE + `",` +
				   ` "osTypeId": "` + TEMPLATE_OS_TYPE_ID + `",` +
				   ` "hypervisor": "` + TEMPLATE_HYPERVISOR + `",` +
				   ` "format": "` + TEMPLATE_FORMAT + `",` +
				   ` "zoneName": "` + TEMPLATE_ZONE_NAME + `",` +
				   ` "projectId": "` + TEMPLATE_PROJECT_ID + `"}`)
}

func TestGetTemplateReturnTemplateIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	templateService := TemplateApi{
		entityService: mockEntityService,
	}

	expectedTemplate := Template{Id: TEMPLATE_ID,
								 Name: TEMPLATE_NAME,
								 Description: TEMPLATE_DESCRIPTION,
								 Size: TEMPLATE_SIZE,
								 IsPublic: TEMPLATE_IS_PUBLIC,
								 IsReady: TEMPLATE_IS_READY,
								 SSHKeyEnabled: TEMPLATE_SSH_KEY_ENABLED,
								 Extractable: TEMPLATE_EXTRACTABLE,
								 OSType: TEMPLATE_OS_TYPE,
								 OSTypeId: TEMPLATE_OS_TYPE_ID,
								 Hypervisor: TEMPLATE_HYPERVISOR,
								 Format: TEMPLATE_FORMAT,
								 ZoneName: TEMPLATE_ZONE_NAME,
								 ProjectId: TEMPLATE_PROJECT_ID,
								}

	mockEntityService.EXPECT().Get(TEMPLATE_ID, gomock.Any()).Return(buildGetTemplateSuccessResponse(), nil)

	//when
	template, _ := templateService.Get(TEMPLATE_ID)

	//then
	if assert.NotNil(t, template) {
		assert.Equal(t, expectedTemplate, *template)
	}
}

func TestGetTemplateReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	templateService := TemplateApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(TEMPLATE_ID, gomock.Any()).Return(nil, mockError)

	//when
	template, err := templateService.Get(TEMPLATE_ID)

	//then
	assert.Nil(t, template)
	assert.Equal(t, mockError, err)

}

func buildListTemplateSuccessResponse() []byte {
	template1 := `{"id":"list_id_1",` +
				   ` "name": "list_name_1",` +
				   ` "description": "list_description_1",` +
				   ` "size": 1,` +
				   ` "isPublic": true,` +
				   ` "isReady": true,` +
				   ` "sshKeyEnabled": true,` +
				   ` "extractable": true,` +
				   ` "osType": "list_os_type_1",` +
				   ` "osTypeId": "list_os_type_id_1",` +
				   ` "hypervisor": "list_hypervisor_1",` +
				   ` "format": "list_format_1",` +
				   ` "zoneName": "list_zone_name_1",` +
				   ` "projectId": "list_project_id_1"}`
	template2 := `{"id":"list_id_2",` +
				   ` "name": "list_name_2",` +
				   ` "description": "list_description_2",` +
				   ` "size": 2,` +
				   ` "isPublic": false,` +
				   ` "isReady": false,` +
				   ` "sshKeyEnabled": false,` +
				   ` "extractable": false,` +
				   ` "osType": "list_os_type_2",` +
				   ` "osTypeId": "list_os_type_id_2",` +
				   ` "hypervisor": "list_hypervisor_2",` +
				   ` "format": "list_format_2",` +
				   ` "zoneName": "list_zone_name_2",` +
				   ` "projectId": "list_project_id_2"}`
	return []byte(`[` + template1 + `,` + template2 + `]`)
}

func TestListTemplateReturnTemplatesIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	templateService := TemplateApi{
		entityService: mockEntityService,
	}

	template1 := Template{Id: "list_id_1",
						  Name: "list_name_1",
						  Description: "list_description_1",
						  Size: 1,
						  IsPublic: true,
						  IsReady: true,
						  SSHKeyEnabled: true,
						  Extractable: true,
						  OSType: "list_os_type_1",
						  OSTypeId: "list_os_type_id_1",
						  Hypervisor: "list_hypervisor_1",
						  Format: "list_format_1",
						  ZoneName: "list_zone_name_1",
						  ProjectId: "list_project_id_1",
				}
	template2 := Template{Id: "list_id_2",
						  Name: "list_name_2",
						  Description: "list_description_2",
						  Size: 2,
						  IsPublic: false,
						  IsReady: false,
						  SSHKeyEnabled: false,
						  Extractable: false,
						  OSType: "list_os_type_2",
						  OSTypeId: "list_os_type_id_2",
						  Hypervisor: "list_hypervisor_2",
						  Format: "list_format_2",
						  ZoneName: "list_zone_name_2",
						  ProjectId: "list_project_id_2",
				}

	expectedTemplates := []Template{template1, template2}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListTemplateSuccessResponse(), nil)

	//when
	templates, _ := templateService.List()

	//then
	if assert.NotNil(t, templates) {
		assert.Equal(t, expectedTemplates, templates)
	}
}


func TestListTemplateReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	templateService := TemplateApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	templates, err := templateService.List()

	//then
	assert.Nil(t, templates)
	assert.Equal(t, mockError, err)

}