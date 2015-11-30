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
	COMPUTE_OFFERING_ID = "some_id"
	COMPUTE_OFFERING_NAME = "test_compute_offering"
	COMPUTE_OFFERING_MEMORY = 4
	COMPUTE_OFFERING_CPU_NUMBER = 2
)

func buildGetSuccessResponse() []byte {
	return  []byte(`{"id": "` + COMPUTE_OFFERING_ID + 
			`","name":"` + COMPUTE_OFFERING_NAME + 
			`","memory":` + strconv.Itoa(COMPUTE_OFFERING_MEMORY) + 
			`,"cpuNumber": ` + strconv.Itoa(COMPUTE_OFFERING_CPU_NUMBER) + `}`)
}

func TestGetComputeOfferingReturnComputeOfferingIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	computeOfferingService := ComputeOfferingApi{
		entityService: mockEntityService,
	}

	expectedComputeOffering := ComputeOffering{Id: COMPUTE_OFFERING_ID, 
										Name: COMPUTE_OFFERING_NAME, 
										Memory: COMPUTE_OFFERING_MEMORY,
										CpuNumber: COMPUTE_OFFERING_CPU_NUMBER}

	mockEntityService.EXPECT().Get(COMPUTE_OFFERING_ID, gomock.Any()).Return(buildGetSuccessResponse(), nil)

	//when
	computeOffering, _ := computeOfferingService.Get(COMPUTE_OFFERING_ID)

	//then
	if assert.NotNil(t, computeOffering) {
		assert.Equal(t, expectedComputeOffering, *computeOffering)
	}
}

func TestGetComputeOfferingReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	computeOfferingService := ComputeOfferingApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(COMPUTE_OFFERING_ID, gomock.Any()).Return(nil, mockError)

	//when
	computeOffering, err := computeOfferingService.Get(COMPUTE_OFFERING_ID)

	//then
	assert.Nil(t, computeOffering)
	assert.Equal(t, mockError, err)

}

func buildListSuccessResponse() []byte {
	return []byte(`[
		{"id": "list_id_1", "name": "list_name_1", "memory": 1, "cpuNumber": 1},
		{"id": "list_id_2", "name": "list_name_2", "memory": 2, "cpuNumber": 2}
	]`)
}

func TestListComputeOfferingReturnComputeOfferingsIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	computeOfferingService := ComputeOfferingApi{
		entityService: mockEntityService,
	}

	expectedComputeOfferings := []ComputeOffering{
		ComputeOffering{
			Id: "list_id_1",
			Name: "list_name_1",
			Memory: 1,
			CpuNumber: 1,
		},
		ComputeOffering{
			Id: "list_id_2",
			Name: "list_name_2",
			Memory: 2,
			CpuNumber: 2,
		},
	}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListSuccessResponse(), nil)

	//when
	computeOfferings, _ := computeOfferingService.List()

	//then
	if assert.NotNil(t, computeOfferings) {
		assert.Equal(t, expectedComputeOfferings, computeOfferings)
	}
}


func TestListComputeOfferingReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	computeOfferingService := ComputeOfferingApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	computeOfferings, err := computeOfferingService.List()

	//then
	assert.Nil(t, computeOfferings)
	assert.Equal(t, mockError, err)

}