package cloudca

import (
	"github.com/cloud-ca/go-cloudca/mocks"
	"github.com/cloud-ca/go-cloudca/mocks/services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

const (
	TEST_DISK_OFFERING_ID           = "some_id"
	TEST_DISK_OFFERING_NAME         = "test_disk_offering"
	TEST_DISK_OFFERING_GBSIZE       = 50
	TEST_DISK_OFFERING_STORAGE_TIER = "performance"
)

func buildDiskOfferingJsonResponse(diskOffering *DiskOffering) []byte {
	return []byte(`{"id": "` + diskOffering.Id +
		`","name":"` + diskOffering.Name +
		`","gbSize":` + strconv.Itoa(diskOffering.GbSize) +
		`,"storageTier": "` + diskOffering.StorageTier + `"}`)
}

func buildListDiskOfferingJsonResponse(diskOfferings []DiskOffering) []byte {
	resp := `[`
	for i, d := range diskOfferings {
		resp += string(buildDiskOfferingJsonResponse(&d))
		if i != len(diskOfferings)-1 {
			resp += `,`
		}
	}
	resp += `]`
	return []byte(resp)
}

func TestGetDiskOfferingReturnDiskOfferingIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	diskOfferingService := DiskOfferingApi{
		entityService: mockEntityService,
	}

	expectedDiskOffering := DiskOffering{Id: TEST_DISK_OFFERING_ID,
		Name:        TEST_DISK_OFFERING_NAME,
		GbSize:      TEST_DISK_OFFERING_GBSIZE,
		StorageTier: TEST_DISK_OFFERING_STORAGE_TIER}

	mockEntityService.EXPECT().Get(TEST_DISK_OFFERING_ID, gomock.Any()).Return(buildDiskOfferingJsonResponse(&expectedDiskOffering), nil)

	//when
	diskOffering, _ := diskOfferingService.Get(TEST_DISK_OFFERING_ID)

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

	mockEntityService.EXPECT().Get(TEST_DISK_OFFERING_ID, gomock.Any()).Return(nil, mockError)

	//when
	diskOffering, err := diskOfferingService.Get(TEST_DISK_OFFERING_ID)

	//then
	assert.Nil(t, diskOffering)
	assert.Equal(t, mockError, err)

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
			Id:          "list_id_1",
			Name:        "list_name_1",
			GbSize:      51,
			StorageTier: "storage_tier_1",
		},
		DiskOffering{
			Id:          "list_id_2",
			Name:        "list_name_2",
			GbSize:      52,
			StorageTier: "storage_tier_2",
		},
	}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListDiskOfferingJsonResponse(expectedDiskOfferings), nil)

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
