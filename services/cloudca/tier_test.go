package cloudca

import (
	"testing"
	"github.com/cloud-ca/go-cloudca/mocks"
	"github.com/cloud-ca/go-cloudca/mocks/services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	TIER_ID = "test_tier_id"
	TIER_NAME = "test_tier"
)

func buildTierJsonResponse(tier *Tier) []byte {
	return  []byte(`{"id":"` + tier.Id + `",` +
				   ` "name": "` + tier.Name + `"}`)
}

func buildListTierJsonResponse(tiers []Tier) []byte {
	resp := `[`
	for i, t := range tiers {
		resp += string(buildTierJsonResponse(&t))
		if i != len(tiers) - 1 {
			resp += `,`
		}
	}
	resp += `]`
	return []byte(resp)
}

func TestGetTierReturnTierIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	tierService := TierApi{
		entityService: mockEntityService,
	}

	expectedTier := Tier{Id: TIER_ID,
								 Name: TIER_NAME}

	mockEntityService.EXPECT().Get(TIER_ID, gomock.Any()).Return(buildTierJsonResponse(&expectedTier), nil)

	//when
	tier, _ := tierService.Get(TIER_ID)

	//then
	if assert.NotNil(t, tier) {
		assert.Equal(t, expectedTier, *tier)
	}
}

func TestGetTierReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	tierService := TierApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(TIER_ID, gomock.Any()).Return(nil, mockError)

	//when
	tier, err := tierService.Get(TIER_ID)

	//then
	assert.Nil(t, tier)
	assert.Equal(t, mockError, err)

}

func TestListTierReturnTiersIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	tierService := TierApi{
		entityService: mockEntityService,
	}

	expectedTiers := []Tier{
		Tier{
			Id: "list_id_1",
			Name: "list_name_1",
		},
		Tier{
			Id: "list_id_2",
			Name: "list_name_2",
		},
	}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListTierJsonResponse(expectedTiers), nil)

	//when
	tiers, _ := tierService.List()

	//then
	if assert.NotNil(t, tiers) {
		assert.Equal(t, expectedTiers, tiers)
	}
}


func TestListTierReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	tierService := TierApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	tiers, err := tierService.List()

	//then
	assert.Nil(t, tiers)
	assert.Equal(t, mockError, err)

}