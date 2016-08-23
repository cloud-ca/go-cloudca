package cloudca

import (
	"fmt"
	"testing"

	"github.com/cloud-ca/go-cloudca/mocks"
	"github.com/cloud-ca/go-cloudca/mocks/services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const ACL_RULE_TEMPLATE = `{
	"id": "%s",
	"aclId": "6145ea41-010c-41f2-a065-2a3a4e98d09d",
	"ruleNumber": "1",
	"cidr": "0.0.0.0/24",
	"action": "Allow",
	"protocol": "TCP",
	"startPort": "80",
	"endPort": "80",
	"trafficType": "Ingress",
	"state": "Active"
}`

func setupMockForNetworkAclRule(t *testing.T) *services_mocks.MockEntityService {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	return services_mocks.NewMockEntityService(ctrl)
}

func createNetworkAclRuleWithId(id string) *NetworkAclRule {
	return &NetworkAclRule{
		Id:          id,
		AclId:       "6145ea41-010c-41f2-a065-2a3a4e98d09d",
		RuleNumber:  "1",
		Cidr:        "0.0.0.0/24",
		Action:      "Allow",
		Protocol:    "TCP",
		StartPort:   "80",
		EndPort:     "80",
		TrafficType: "Ingress",
		State:       "Active",
	}
}

func TestGetNetworkAclRuleById(t *testing.T) {
	// given
	mockEntityService := setupMockForNetworkAclRule(t)
	networkAclRuleService := NetworkAclRuleApi{
		entityService: mockEntityService,
	}

	expectedId := "rule_0"
	expectedNetworkAclRule := *createNetworkAclRuleWithId(expectedId)

	response := fmt.Sprintf(ACL_RULE_TEMPLATE, expectedId)
	mockEntityService.EXPECT().Get(expectedId, gomock.Any()).Return([]byte(response), nil)

	// when
	networkAclRule, _ := networkAclRuleService.Get(expectedId)

	// then
	assert.Equal(t, expectedNetworkAclRule, *networkAclRule)
}

func TestListNetworkAclRulesWithOptions(t *testing.T) {
	// given
	mockEntityService := setupMockForNetworkAclRule(t)
	networkAclRuleService := NetworkAclRuleApi{
		entityService: mockEntityService,
	}

	id1, id2 := "1234", "4321"
	rule1, rule2 := fmt.Sprintf(ACL_RULE_TEMPLATE, id1), fmt.Sprintf(ACL_RULE_TEMPLATE, id2)
	response := fmt.Sprintf("[ %s, %s ]", rule1, rule2)
	mockEntityService.EXPECT().List(gomock.Any()).Return([]byte(response), nil)

	// when
	rules, _ := networkAclRuleService.ListWithOptions(map[string]string{})

	// then
	assert.Equal(t, id1, rules[0].Id)
	assert.Equal(t, id2, rules[1].Id)
}

func TestCreateNetworkAclRule(t *testing.T) {
	// given
	mockEntityService := setupMockForNetworkAclRule(t)
	networkAclRuleService := NetworkAclRuleApi{
		entityService: mockEntityService,
	}

	expectedId := "adsf"
	response := fmt.Sprintf(ACL_RULE_TEMPLATE, expectedId)
	expectedNetworkAclRule := *createNetworkAclRuleWithId(expectedId)

	mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return([]byte(response), nil)

	// when
	rule, _ := networkAclRuleService.Create(expectedNetworkAclRule)

	// then
	assert.Equal(t, expectedNetworkAclRule, *rule)
}

func TestDeleteNetworkAclRuleReturnsSuccess_ifNoErrorsOccur(t *testing.T) {
	// given
	mockEntityService := setupMockForNetworkAclRule(t)
	networkAclRuleService := NetworkAclRuleApi{
		entityService: mockEntityService,
	}

	expectedId := "id0"
	mockEntityService.EXPECT().Delete(expectedId, gomock.Any(), gomock.Any()).Return([]byte{}, nil)

	// when
	success, _ := networkAclRuleService.Delete(expectedId)

	// then
	assert.True(t, success)
}

func TestDeleteNetworkAclRuleReturnsFailure_ifErrorOccurred(t *testing.T) {
	// given
	mockEntityService := setupMockForNetworkAclRule(t)
	networkAclRuleService := NetworkAclRuleApi{
		entityService: mockEntityService,
	}

	expectedId := "id0"
	mockEntityService.EXPECT().Delete(expectedId, gomock.Any(), gomock.Any()).Return(nil, mocks.MockError{Message: "asdf"})

	// when
	success, _ := networkAclRuleService.Delete(expectedId)

	// then
	assert.False(t, success)
}
