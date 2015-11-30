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
	TEST_TIER_ID = "test_tier_id"
	TEST_TIER_NAME = "test_tier"
	TEST_TIER_ZONE_ID = "test_tier_zone_id"
	TEST_TIER_ZONE_NAME = "test_tier_zone_name"
	TEST_TIER_CIDR = "test_tier_cidr"
	TEST_TIER_TYPE = "test_tier_type"
	TEST_TIER_STATE = "test_tier_state"
	TEST_TIER_GATEWAY = "test_tier_gateway"
	TEST_TIER_NETWORK_OFFERING_ID = "test_tier_network_offering_id"
	TEST_TIER_IS_SYSTEM = false
	TEST_TIER_VPC_ID = "test_tier_vpc_id"
	TEST_TIER_DOMAIN = "test_tier_domain"
	TEST_TIER_DOMAIN_ID = "test_tier_domain_id"
	TEST_TIER_PROJECT = "test_tier_project"
	TEST_TIER_PROJECT_ID = "test_tier_project_id"
	TEST_TIER_ACL_ID = "test_tier_acl_id"
)

func buildTestTierJsonResponse(tier *Tier) []byte {
	return  []byte(`{"id":"` + tier.Id + `",` +
				   ` "name":"` + tier.Name + `",` +
				   ` "zoneid":"` + tier.ZoneId + `",` +
				   ` "zonename":"` + tier.ZoneName + `",` +
				   ` "cidr":"` + tier.Cidr + `",` +
				   ` "type":"` + tier.Type + `",` +
				   ` "state":"` + tier.State + `",` +
				   ` "gateway":"` + tier.Gateway + `",` +
				   ` "networkofferingid":"` + tier.NetworkOfferingId + `",` +
				   ` "issystem":` + strconv.FormatBool(tier.IsSystem) + `,` +
				   ` "vpcid":"` + tier.VpcId + `",` +
				   ` "domain":"` + tier.Domain + `",` +
				   ` "domainid":"` + tier.DomainId + `",` +
				   ` "project":"` + tier.Project + `",` +
				   ` "projectid":"` + tier.ProjectId + `",` +
				   ` "aclId":"` + tier.AclId + `"}`)
}

func buildListTestTierJsonResponse(tiers []Tier) []byte {
	resp := `[`
	for i, t := range tiers {
		resp += string(buildTestTierJsonResponse(&t))
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

	expectedTier := Tier{Id: TEST_TIER_ID,
						 Name: TEST_TIER_NAME,
						 ZoneId: TEST_TIER_ZONE_ID,
						 ZoneName: TEST_TIER_ZONE_NAME,
						 Cidr: TEST_TIER_CIDR,
						 Type: TEST_TIER_TYPE,
						 Gateway: TEST_TIER_GATEWAY,
						 NetworkOfferingId: TEST_TIER_NETWORK_OFFERING_ID,
						 IsSystem: TEST_TIER_IS_SYSTEM,
						 VpcId: TEST_TIER_VPC_ID,
						 Domain: TEST_TIER_DOMAIN,
						 DomainId: TEST_TIER_DOMAIN_ID,
						 Project: TEST_TIER_PROJECT,
						 ProjectId: TEST_TIER_PROJECT_ID,
						 AclId: TEST_TIER_ACL_ID}

	mockEntityService.EXPECT().Get(TEST_TIER_ID, gomock.Any()).Return(buildTestTierJsonResponse(&expectedTier), nil)

	//when
	tier, _ := tierService.Get(TEST_TIER_ID)

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

	mockEntityService.EXPECT().Get(TEST_TIER_ID, gomock.Any()).Return(nil, mockError)

	//when
	tier, err := tierService.Get(TEST_TIER_ID)

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

	expectedTier1 := Tier{Id: "list_id_1",
						 Name: "list_name_1",
						 ZoneId: "list_zone_id_1",
						 ZoneName: "list_zone_name_1",
						 Cidr: "list_cidr_1",
						 Type: "list_type_1",
						 Gateway: "list_gateway_1",
						 NetworkOfferingId: "list_network_offering_id_1",
						 IsSystem: true,
						 VpcId: "list_vpc_id_1",
						 Domain: "list_domain_1",
						 DomainId: "list_domain_id_1",
						 Project: "list_project_1",
						 ProjectId: "list_project_id_1",
						 AclId: "list_acl_id_1"}

	expectedTier2 := Tier{Id: "list_id_2",
						 Name: "list_name_2",
						 ZoneId: "list_zone_id_2",
						 ZoneName: "list_zone_name_2",
						 Cidr: "list_cidr_2",
						 Type: "list_type_2",
						 Gateway: "list_gateway_2",
						 NetworkOfferingId: "list_network_offering_id_2",
						 IsSystem: false,
						 VpcId: "list_vpc_id_2",
						 Domain: "list_domain_2",
						 DomainId: "list_domain_id_2",
						 Project: "list_project_2",
						 ProjectId: "list_project_id_2",
						 AclId: "list_acl_id_2"}

	expectedTiers := []Tier{expectedTier1, expectedTier2}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListTestTierJsonResponse(expectedTiers), nil)

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