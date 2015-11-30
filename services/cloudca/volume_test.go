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
	VOLUME_ID = "test_volume_id"
	VOLUME_NAME = "test_volume"
	VOLUME_TYPE = "test_volume_type"
	VOLUME_CREATION_DATE = "test_volume_creation_date"
	VOLUME_SIZE = 500
	VOLUME_DISK_OFFERING_ID = "test_volume_disk_offering_id"
	VOLUME_TEMPLATE_ID = "test_volume_template_id"
	VOLUME_STORAGE_TIER = "test_volume_storage_tier"
	VOLUME_ZONE_NAME = "test_volume_zone_name"
	VOLUME_STATE = "test_volume_state"
	VOLUME_INSTANCE_NAME = "test_volume_instance_name"
	VOLUME_INSTANCE_ID = "test_volume_instance_id"
	VOLUME_INSTANCE_STATE = "test_volume_instance_state"
)

func buildVolumeJsonResponse(volume *Volume) []byte {
	return  []byte(`{"id":"` + volume.Id + `",` +
				   ` "name": "` + volume.Name + `",` +
				   ` "type": "` + volume.Type + `",` +
				   ` "creationDate": "` + volume.CreationDate + `",` +
				   ` "size": ` + strconv.Itoa(volume.Size) + `,` +
				   ` "diskOfferingId": "` + volume.DiskOfferingId + `",` +
				   ` "templateId": "` + volume.TemplateId + `",` +
				   ` "storageTier": "` + volume.StorageTier + `",` +
				   ` "zoneName": "` + volume.ZoneName + `",` +
				   ` "state": "` + volume.State + `",` +
				   ` "instanceName": "` + volume.InstanceName + `",` +
				   ` "instanceId": "` + volume.InstanceId + `",` +
				   ` "instanceState": "` + volume.InstanceState + `"}`)
}

func buildListVolumeJsonResponse(volumes []Volume) []byte {
	resp := `[`
	for i, v := range volumes {
		resp += string(buildVolumeJsonResponse(&v))
		if i != len(volumes) - 1 {
			resp += `,`
		}
	}
	resp += `]`
	return []byte(resp)
}

func TestGetVolumeReturnVolumeIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	volumeService := VolumeApi{
		entityService: mockEntityService,
	}

	expectedVolume := Volume{Id: VOLUME_ID,
						   	 Name: VOLUME_NAME,
						   	 Type: VOLUME_TYPE,
						   	 CreationDate: VOLUME_CREATION_DATE,
						   	 Size: VOLUME_SIZE,
						   	 DiskOfferingId: VOLUME_DISK_OFFERING_ID,
						   	 TemplateId: VOLUME_TEMPLATE_ID,
						   	 StorageTier: VOLUME_STORAGE_TIER,
						   	 ZoneName: VOLUME_ZONE_NAME,
						   	 State: VOLUME_STATE,
						   	 InstanceName: VOLUME_INSTANCE_NAME,
						   	 InstanceId: VOLUME_INSTANCE_ID,
						   	 InstanceState: VOLUME_INSTANCE_STATE}

	mockEntityService.EXPECT().Get(VOLUME_ID, gomock.Any()).Return(buildVolumeJsonResponse(&expectedVolume), nil)

	//when
	volume, _ := volumeService.Get(VOLUME_ID)

	//then
	if assert.NotNil(t, volume) {
		assert.Equal(t, expectedVolume, *volume)
	}
}

func TestGetVolumeReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	volumeService := VolumeApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(VOLUME_ID, gomock.Any()).Return(nil, mockError)

	//when
	volume, err := volumeService.Get(VOLUME_ID)

	//then
	assert.Nil(t, volume)
	assert.Equal(t, mockError, err)

}

func TestListVolumeReturnVolumesIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	volumeService := VolumeApi{
		entityService: mockEntityService,
	}

	expectedVolume1 := Volume{Id: "list_id_1",
						   	 Name: "list_name_1",
						   	 Type: "list_type_1",
						   	 CreationDate: "list_creation_date_1",
						   	 Size: 1215,
						   	 DiskOfferingId: "list_disk_offering_id_1",
						   	 TemplateId: "list_template_id_1",
						   	 StorageTier: "list_storage_tier_1",
						   	 ZoneName: "list_zone_name_1",
						   	 State: "list_state_1",
						   	 InstanceName: "list_instance_name_1",
						   	 InstanceId: "list_instance_id_1",
						   	 InstanceState: "list_instance_state_1"}
	expectedVolume2 := Volume{Id: "list_id_2",
						   	 Name: "list_name_2",
						   	 Type: "list_type_2",
						   	 CreationDate: "list_creation_date_2",
						   	 Size: 54582,
						   	 DiskOfferingId: "list_disk_offering_id_2",
						   	 TemplateId: "list_template_id_2",
						   	 StorageTier: "list_storage_tier_2",
						   	 ZoneName: "list_zone_name_2",
						   	 State: "list_state_2",
						   	 InstanceName: "list_instance_name_2",
						   	 InstanceId: "list_instance_id_2",
						   	 InstanceState: "list_instance_state_2"}

	expectedVolumes := []Volume{expectedVolume1, expectedVolume2}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListVolumeJsonResponse(expectedVolumes), nil)

	//when
	volumes, _ := volumeService.List()

	//then
	if assert.NotNil(t, volumes) {
		assert.Equal(t, expectedVolumes, volumes)
	}
}


func TestListVolumeReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	volumeService := VolumeApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	volumes, err := volumeService.List()

	//then
	assert.Nil(t, volumes)
	assert.Equal(t, mockError, err)

}