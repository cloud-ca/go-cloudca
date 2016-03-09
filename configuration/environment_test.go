package configuration

import (
  "testing"
  // "github.com/cloud-ca/go-cloudca/mocks/api_mocks"
  "github.com/golang/mock/gomock"
  "github.com/stretchr/testify/assert"
)

const (
  TEST_ENVIRONMENT_ID = "test_environment_id"
  TEST_ENVIRONMENT_NAME = "test_environment_name"
  TEST_ENVIRONMENT_DESCRIPTION = "test_environment_description"
)

func buildTestEnvironmentJsonResponse(environment *Environment) []byte {
  return  []byte(`{"id": "` + environment.Id + `", ` +
      `"name":"` + environment.Name + `", ` +
      `"description":"` + environment.Description + `"}`)
}

func buildListTestEnvironmentJsonResponse(environments []Environment) []byte {
  resp := `[`
  for i, env := range environments {
    resp += string(buildTestEnvironmentJsonResponse(&env))
    if i != len(environments) - 1 {
      resp += `,`
    }
  }
  resp += `]`
  return []byte(resp)
}

func TestConfigurationType(t *testing.T){
  ctrl := gomock.NewController(t)
  defer ctrl.Finish()
  
  // this is a weak test. It validates that the constant is the correct API endpoint,
  // because the buildUrl method is unaccessible
  assert.Equal(t, ENVIRONMENT_CONFIGURATION_TYPE, "environments")
}

func TestGetEnvironmentReturnEnvironmentIfSuccess(t *testing.T) {
  //given
  ctrl := gomock.NewController(t)
  defer ctrl.Finish()

  mockConfigurationService := configuration_mocks.NewMockConfigurationService(ctrl)
  environmentService := EnvironmentApi{
    configurationService: mockConfigurationService,
  }

  expectedEnvironment := Environment{Id: "envId", 
                 Name: "envName",
                 Description: "testDesc",
               }




  mockEntityService.EXPECT().Get(TEST_ENVIRONMENT_ID, gomock.Any()).Return(buildTestEnvironmentJsonResponse(&expectedEnvironment), nil)

  //when
  environment, _ := environmentService.Get(TEST_ENVIRONMENT_ID)

  //then
  if assert.NotNil(t, environment) {
    assert.Equal(t, expectedEnvironment, *environment)
  }
}

// func TestGetEnvironmentReturnNilWithErrorIfError(t *testing.T) {
//   //given
//   ctrl := gomock.NewController(t)
//   defer ctrl.Finish()

//   mockEntityService := services_mocks.NewMockEntityService(ctrl)

//   environmentService := EnvironmentApi{
//     entityService: mockEntityService,
//   }

//   mockError := mocks.MockError{"some_get_error"}

//   mockEntityService.EXPECT().Get(TEST_ENVIRONMENT_ID, gomock.Any()).Return(nil, mockError)

//   //when
//   environment, err := environmentService.Get(TEST_ENVIRONMENT_ID)

//   //then
//   assert.Nil(t, environment)
//   assert.Equal(t, mockError, err)

// }

// func TestListEnvironmentReturnEnvironmentsIfSuccess(t *testing.T) {
//   //given
//   ctrl := gomock.NewController(t)
//   defer ctrl.Finish()

//   mockEntityService := services_mocks.NewMockEntityService(ctrl)

//   environmentService := EnvironmentApi{
//     entityService: mockEntityService,
//   }

//   expectedEnvironment1 := Environment{Id: "list_id_1", 
//                  Name: "list_name_1",
//                  State: "list_state_1",
//                  TemplateId: "list_template_id_1",
//                  TemplateName: "list_template_name_1",
//                  IsPasswordEnabled: false,
//                  IsSSHKeyEnabled: true,
//                  Username: "list_username_1",
//                  ComputeOfferingId: "list_compute_offering_id_1",
//                  ComputeOfferingName: "list_compute_offering_name_1",
//                  CpuCount: 2,
//                  MemoryInMB: 12425,
//                  ZoneId: "list_zone_id_1",
//                  ZoneName: "list_zone_name_1",
//                  ProjectId: "list_project_id_1",
//                  NetworkId: "list_network_id_1",
//                  NetworkName: "list_network_name_1",
//                  MacAddress: "list_mac_address_1",
//                  VolumeIdToAttach: "list_volume_id_to_attach_1",
//                  IpAddress: "list_ip_address_1",
//                  PublicKey: "list_public_key_1",
//                  UserData: "list_user_data_1"}

//   expectedEnvironments := []Environment{expectedEnvironment1}

//   mockEntityService.EXPECT().List(gomock.Any()).Return(buildListTestEnvironmentJsonResponse(expectedEnvironments), nil)

//   //when
//   environments, _ := environmentService.List()

//   //then
//   if assert.NotNil(t, environments) {
//     assert.Equal(t, expectedEnvironments, environments)
//   }
// }


// func TestListEnvironmentReturnNilWithErrorIfError(t *testing.T) {
//   //given
//   ctrl := gomock.NewController(t)
//   defer ctrl.Finish()

//   mockEntityService := services_mocks.NewMockEntityService(ctrl)

//   environmentService := EnvironmentApi{
//     entityService: mockEntityService,
//   }

//   mockError := mocks.MockError{"some_list_error"}

//   mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

//   //when
//   environments, err := environmentService.List()

//   //then
//   assert.Nil(t, environments)
//   assert.Equal(t, mockError, err)

// }

// func TestCreateEnvironmentReturnCreatedEnvironmentIfSuccess(t *testing.T) {
//   //given
//   ctrl := gomock.NewController(t)
//   defer ctrl.Finish()

//   mockEntityService := services_mocks.NewMockEntityService(ctrl)

//   environmentService := EnvironmentApi{
//     entityService: mockEntityService,
//   }

//   environmentToCreate := Environment{Id: "new_id", 
//                  Name: "new_name",
//                  TemplateId: "templateId",
//                  ComputeOfferingId: "computeOfferingId",
//                  NetworkId: "networkId"}

//   mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return([]byte(`{"id":"new_id", "password": "new_password"}`), nil)

//   //when
//   createdEnvironment, _ := environmentService.Create(environmentToCreate)

//   //then
//   if assert.NotNil(t, createdEnvironment) {
//     assert.Equal(t, "new_password", createdEnvironment.Password)
//   }
// }

// func TestCreateEnvironmentReturnNilWithErrorIfError(t *testing.T) {
//   //given
//   ctrl := gomock.NewController(t)
//   defer ctrl.Finish()

//   mockEntityService := services_mocks.NewMockEntityService(ctrl)

//   environmentService := EnvironmentApi{
//     entityService: mockEntityService,
//   }

//   mockError := mocks.MockError{"some_create_environment_error"}

//   mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, mockError)

//   environmentToCreate := Environment{Name: "new_name",
//                  TemplateId: "templateId",
//                  ComputeOfferingId: "computeOfferingId",
//                  NetworkId: "networkId"}

//   //when
//   createdEnvironment, err := environmentService.Create(environmentToCreate)

//   //then
//   assert.Nil(t, createdEnvironment)
//   assert.Equal(t, mockError, err)

// }