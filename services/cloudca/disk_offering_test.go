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
	DISK_OFFERING_ID = "some_id"
	DISK_OFFERING_NAME = "test_disk_offering"
	DISK_OFFERING_GBSIZE = 50
	DISK_OFFERING_STORAGE_TIER = "performance"
)

func buildGetDiskOfferingSuccessResponse() []byte {
	return  []byte(`{"id": "` + DISK_OFFERING_ID + 
			`","name":"` + DISK_OFFERING_NAME + 
			`","gbSize":` + strconv.Itoa(DISK_OFFERING_GBSIZE) + 
			`,"storageTier": "` + DISK_OFFERING_STORAGE_TIER + `"}`)
}

func TestGetDiskOfferingReturnDiskOfferingIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	diskOfferingService := DiskOfferingApi{
		entityService: mockEntityService,
	}

	expectedDiskOffering := DiskOffering{Id: DISK_OFFERING_ID, 
										Name: DISK_OFFERING_NAME, 
										GbSize: DISK_OFFERING_GBSIZE,
										StorageTier: DISK_OFFERING_STORAGE_TIER}

	mockEntityService.EXPECT().Get(DISK_OFFERING_ID, gomock.Any()).Return(buildGetDiskOfferingSuccessResponse(), nil)

	//when
	diskOffering, _ := diskOfferingService.Get(DISK_OFFERING_ID)

	//then
	if assert.NotNil(t, diskOffering) {
		assert.Equal(t, expectedDiskOffering, *diskOffering)
	}
}

func TestGetDiskOfferingReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	diskOfferingService := DiskOfferingApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(DISK_OFFERING_ID, gomock.Any()).Return(nil, mockError)

	//when
	diskOffering, err := diskOfferingService.Get(DISK_OFFERING_ID)

	//then
	assert.Nil(t, diskOffering)
	assert.Equal(t, mockError, err)

}

func buildListDiskOfferingSuccessResponse() []byte {
	return []byte(`[
		{"id": "list_id_1", "name": "list_name_1", "gbSize": 51, "storageTier": "storage_tier_1"},
		{"id": "list_id_2", "name": "list_name_2", "gbSize": 52, "storageTier": "storage_tier_2"}
	]`)
}

func TestListDiskOfferingReturnDiskOfferingsIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	diskOfferingService := DiskOfferingApi{
		entityService: mockEntityService,
	}

	expectedDiskOfferings := []DiskOffering{
		DiskOffering{
			Id: "list_id_1",
			Name: "list_name_1",
			GbSize: 51,
			StorageTier: "storage_tier_1",
		},
		DiskOffering{
			Id: "list_id_2",
			Name: "list_name_2",
			GbSize: 52,
			StorageTier: "storage_tier_2",
		},
	}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListDiskOfferingSuccessResponse(), nil)

	//when
	diskOfferings, _ := diskOfferingService.List()

	//then
	if assert.NotNil(t, diskOfferings) {
		assert.Equal(t, expectedDiskOfferings, diskOfferings)
	}
}


func TestListDiskOfferingReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	diskOfferingService := DiskOfferingApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	diskOfferings, err := diskOfferingService.List()

	//then
	assert.Nil(t, diskOfferings)
	assert.Equal(t, mockError, err)

}