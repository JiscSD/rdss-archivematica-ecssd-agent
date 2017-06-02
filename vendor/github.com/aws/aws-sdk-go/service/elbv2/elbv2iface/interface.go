// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package elbv2iface provides an interface to enable mocking the Elastic Load Balancing service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package elbv2iface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

// ELBV2API provides an interface to enable mocking the
// elbv2.ELBV2 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Elastic Load Balancing.
//    func myFunc(svc elbv2iface.ELBV2API) bool {
//        // Make svc.AddTags request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := elbv2.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockELBV2Client struct {
//        elbv2iface.ELBV2API
//    }
//    func (m *mockELBV2Client) AddTags(input *elbv2.AddTagsInput) (*elbv2.AddTagsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockELBV2Client{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ELBV2API interface {
	AddTags(*elbv2.AddTagsInput) (*elbv2.AddTagsOutput, error)
	AddTagsWithContext(aws.Context, *elbv2.AddTagsInput, ...request.Option) (*elbv2.AddTagsOutput, error)
	AddTagsRequest(*elbv2.AddTagsInput) (*request.Request, *elbv2.AddTagsOutput)

	CreateListener(*elbv2.CreateListenerInput) (*elbv2.CreateListenerOutput, error)
	CreateListenerWithContext(aws.Context, *elbv2.CreateListenerInput, ...request.Option) (*elbv2.CreateListenerOutput, error)
	CreateListenerRequest(*elbv2.CreateListenerInput) (*request.Request, *elbv2.CreateListenerOutput)

	CreateLoadBalancer(*elbv2.CreateLoadBalancerInput) (*elbv2.CreateLoadBalancerOutput, error)
	CreateLoadBalancerWithContext(aws.Context, *elbv2.CreateLoadBalancerInput, ...request.Option) (*elbv2.CreateLoadBalancerOutput, error)
	CreateLoadBalancerRequest(*elbv2.CreateLoadBalancerInput) (*request.Request, *elbv2.CreateLoadBalancerOutput)

	CreateRule(*elbv2.CreateRuleInput) (*elbv2.CreateRuleOutput, error)
	CreateRuleWithContext(aws.Context, *elbv2.CreateRuleInput, ...request.Option) (*elbv2.CreateRuleOutput, error)
	CreateRuleRequest(*elbv2.CreateRuleInput) (*request.Request, *elbv2.CreateRuleOutput)

	CreateTargetGroup(*elbv2.CreateTargetGroupInput) (*elbv2.CreateTargetGroupOutput, error)
	CreateTargetGroupWithContext(aws.Context, *elbv2.CreateTargetGroupInput, ...request.Option) (*elbv2.CreateTargetGroupOutput, error)
	CreateTargetGroupRequest(*elbv2.CreateTargetGroupInput) (*request.Request, *elbv2.CreateTargetGroupOutput)

	DeleteListener(*elbv2.DeleteListenerInput) (*elbv2.DeleteListenerOutput, error)
	DeleteListenerWithContext(aws.Context, *elbv2.DeleteListenerInput, ...request.Option) (*elbv2.DeleteListenerOutput, error)
	DeleteListenerRequest(*elbv2.DeleteListenerInput) (*request.Request, *elbv2.DeleteListenerOutput)

	DeleteLoadBalancer(*elbv2.DeleteLoadBalancerInput) (*elbv2.DeleteLoadBalancerOutput, error)
	DeleteLoadBalancerWithContext(aws.Context, *elbv2.DeleteLoadBalancerInput, ...request.Option) (*elbv2.DeleteLoadBalancerOutput, error)
	DeleteLoadBalancerRequest(*elbv2.DeleteLoadBalancerInput) (*request.Request, *elbv2.DeleteLoadBalancerOutput)

	DeleteRule(*elbv2.DeleteRuleInput) (*elbv2.DeleteRuleOutput, error)
	DeleteRuleWithContext(aws.Context, *elbv2.DeleteRuleInput, ...request.Option) (*elbv2.DeleteRuleOutput, error)
	DeleteRuleRequest(*elbv2.DeleteRuleInput) (*request.Request, *elbv2.DeleteRuleOutput)

	DeleteTargetGroup(*elbv2.DeleteTargetGroupInput) (*elbv2.DeleteTargetGroupOutput, error)
	DeleteTargetGroupWithContext(aws.Context, *elbv2.DeleteTargetGroupInput, ...request.Option) (*elbv2.DeleteTargetGroupOutput, error)
	DeleteTargetGroupRequest(*elbv2.DeleteTargetGroupInput) (*request.Request, *elbv2.DeleteTargetGroupOutput)

	DeregisterTargets(*elbv2.DeregisterTargetsInput) (*elbv2.DeregisterTargetsOutput, error)
	DeregisterTargetsWithContext(aws.Context, *elbv2.DeregisterTargetsInput, ...request.Option) (*elbv2.DeregisterTargetsOutput, error)
	DeregisterTargetsRequest(*elbv2.DeregisterTargetsInput) (*request.Request, *elbv2.DeregisterTargetsOutput)

	DescribeAccountLimits(*elbv2.DescribeAccountLimitsInput) (*elbv2.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsWithContext(aws.Context, *elbv2.DescribeAccountLimitsInput, ...request.Option) (*elbv2.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsRequest(*elbv2.DescribeAccountLimitsInput) (*request.Request, *elbv2.DescribeAccountLimitsOutput)

	DescribeListeners(*elbv2.DescribeListenersInput) (*elbv2.DescribeListenersOutput, error)
	DescribeListenersWithContext(aws.Context, *elbv2.DescribeListenersInput, ...request.Option) (*elbv2.DescribeListenersOutput, error)
	DescribeListenersRequest(*elbv2.DescribeListenersInput) (*request.Request, *elbv2.DescribeListenersOutput)

	DescribeListenersPages(*elbv2.DescribeListenersInput, func(*elbv2.DescribeListenersOutput, bool) bool) error
	DescribeListenersPagesWithContext(aws.Context, *elbv2.DescribeListenersInput, func(*elbv2.DescribeListenersOutput, bool) bool, ...request.Option) error

	DescribeLoadBalancerAttributes(*elbv2.DescribeLoadBalancerAttributesInput) (*elbv2.DescribeLoadBalancerAttributesOutput, error)
	DescribeLoadBalancerAttributesWithContext(aws.Context, *elbv2.DescribeLoadBalancerAttributesInput, ...request.Option) (*elbv2.DescribeLoadBalancerAttributesOutput, error)
	DescribeLoadBalancerAttributesRequest(*elbv2.DescribeLoadBalancerAttributesInput) (*request.Request, *elbv2.DescribeLoadBalancerAttributesOutput)

	DescribeLoadBalancers(*elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancersWithContext(aws.Context, *elbv2.DescribeLoadBalancersInput, ...request.Option) (*elbv2.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancersRequest(*elbv2.DescribeLoadBalancersInput) (*request.Request, *elbv2.DescribeLoadBalancersOutput)

	DescribeLoadBalancersPages(*elbv2.DescribeLoadBalancersInput, func(*elbv2.DescribeLoadBalancersOutput, bool) bool) error
	DescribeLoadBalancersPagesWithContext(aws.Context, *elbv2.DescribeLoadBalancersInput, func(*elbv2.DescribeLoadBalancersOutput, bool) bool, ...request.Option) error

	DescribeRules(*elbv2.DescribeRulesInput) (*elbv2.DescribeRulesOutput, error)
	DescribeRulesWithContext(aws.Context, *elbv2.DescribeRulesInput, ...request.Option) (*elbv2.DescribeRulesOutput, error)
	DescribeRulesRequest(*elbv2.DescribeRulesInput) (*request.Request, *elbv2.DescribeRulesOutput)

	DescribeSSLPolicies(*elbv2.DescribeSSLPoliciesInput) (*elbv2.DescribeSSLPoliciesOutput, error)
	DescribeSSLPoliciesWithContext(aws.Context, *elbv2.DescribeSSLPoliciesInput, ...request.Option) (*elbv2.DescribeSSLPoliciesOutput, error)
	DescribeSSLPoliciesRequest(*elbv2.DescribeSSLPoliciesInput) (*request.Request, *elbv2.DescribeSSLPoliciesOutput)

	DescribeTags(*elbv2.DescribeTagsInput) (*elbv2.DescribeTagsOutput, error)
	DescribeTagsWithContext(aws.Context, *elbv2.DescribeTagsInput, ...request.Option) (*elbv2.DescribeTagsOutput, error)
	DescribeTagsRequest(*elbv2.DescribeTagsInput) (*request.Request, *elbv2.DescribeTagsOutput)

	DescribeTargetGroupAttributes(*elbv2.DescribeTargetGroupAttributesInput) (*elbv2.DescribeTargetGroupAttributesOutput, error)
	DescribeTargetGroupAttributesWithContext(aws.Context, *elbv2.DescribeTargetGroupAttributesInput, ...request.Option) (*elbv2.DescribeTargetGroupAttributesOutput, error)
	DescribeTargetGroupAttributesRequest(*elbv2.DescribeTargetGroupAttributesInput) (*request.Request, *elbv2.DescribeTargetGroupAttributesOutput)

	DescribeTargetGroups(*elbv2.DescribeTargetGroupsInput) (*elbv2.DescribeTargetGroupsOutput, error)
	DescribeTargetGroupsWithContext(aws.Context, *elbv2.DescribeTargetGroupsInput, ...request.Option) (*elbv2.DescribeTargetGroupsOutput, error)
	DescribeTargetGroupsRequest(*elbv2.DescribeTargetGroupsInput) (*request.Request, *elbv2.DescribeTargetGroupsOutput)

	DescribeTargetGroupsPages(*elbv2.DescribeTargetGroupsInput, func(*elbv2.DescribeTargetGroupsOutput, bool) bool) error
	DescribeTargetGroupsPagesWithContext(aws.Context, *elbv2.DescribeTargetGroupsInput, func(*elbv2.DescribeTargetGroupsOutput, bool) bool, ...request.Option) error

	DescribeTargetHealth(*elbv2.DescribeTargetHealthInput) (*elbv2.DescribeTargetHealthOutput, error)
	DescribeTargetHealthWithContext(aws.Context, *elbv2.DescribeTargetHealthInput, ...request.Option) (*elbv2.DescribeTargetHealthOutput, error)
	DescribeTargetHealthRequest(*elbv2.DescribeTargetHealthInput) (*request.Request, *elbv2.DescribeTargetHealthOutput)

	ModifyListener(*elbv2.ModifyListenerInput) (*elbv2.ModifyListenerOutput, error)
	ModifyListenerWithContext(aws.Context, *elbv2.ModifyListenerInput, ...request.Option) (*elbv2.ModifyListenerOutput, error)
	ModifyListenerRequest(*elbv2.ModifyListenerInput) (*request.Request, *elbv2.ModifyListenerOutput)

	ModifyLoadBalancerAttributes(*elbv2.ModifyLoadBalancerAttributesInput) (*elbv2.ModifyLoadBalancerAttributesOutput, error)
	ModifyLoadBalancerAttributesWithContext(aws.Context, *elbv2.ModifyLoadBalancerAttributesInput, ...request.Option) (*elbv2.ModifyLoadBalancerAttributesOutput, error)
	ModifyLoadBalancerAttributesRequest(*elbv2.ModifyLoadBalancerAttributesInput) (*request.Request, *elbv2.ModifyLoadBalancerAttributesOutput)

	ModifyRule(*elbv2.ModifyRuleInput) (*elbv2.ModifyRuleOutput, error)
	ModifyRuleWithContext(aws.Context, *elbv2.ModifyRuleInput, ...request.Option) (*elbv2.ModifyRuleOutput, error)
	ModifyRuleRequest(*elbv2.ModifyRuleInput) (*request.Request, *elbv2.ModifyRuleOutput)

	ModifyTargetGroup(*elbv2.ModifyTargetGroupInput) (*elbv2.ModifyTargetGroupOutput, error)
	ModifyTargetGroupWithContext(aws.Context, *elbv2.ModifyTargetGroupInput, ...request.Option) (*elbv2.ModifyTargetGroupOutput, error)
	ModifyTargetGroupRequest(*elbv2.ModifyTargetGroupInput) (*request.Request, *elbv2.ModifyTargetGroupOutput)

	ModifyTargetGroupAttributes(*elbv2.ModifyTargetGroupAttributesInput) (*elbv2.ModifyTargetGroupAttributesOutput, error)
	ModifyTargetGroupAttributesWithContext(aws.Context, *elbv2.ModifyTargetGroupAttributesInput, ...request.Option) (*elbv2.ModifyTargetGroupAttributesOutput, error)
	ModifyTargetGroupAttributesRequest(*elbv2.ModifyTargetGroupAttributesInput) (*request.Request, *elbv2.ModifyTargetGroupAttributesOutput)

	RegisterTargets(*elbv2.RegisterTargetsInput) (*elbv2.RegisterTargetsOutput, error)
	RegisterTargetsWithContext(aws.Context, *elbv2.RegisterTargetsInput, ...request.Option) (*elbv2.RegisterTargetsOutput, error)
	RegisterTargetsRequest(*elbv2.RegisterTargetsInput) (*request.Request, *elbv2.RegisterTargetsOutput)

	RemoveTags(*elbv2.RemoveTagsInput) (*elbv2.RemoveTagsOutput, error)
	RemoveTagsWithContext(aws.Context, *elbv2.RemoveTagsInput, ...request.Option) (*elbv2.RemoveTagsOutput, error)
	RemoveTagsRequest(*elbv2.RemoveTagsInput) (*request.Request, *elbv2.RemoveTagsOutput)

	SetIpAddressType(*elbv2.SetIpAddressTypeInput) (*elbv2.SetIpAddressTypeOutput, error)
	SetIpAddressTypeWithContext(aws.Context, *elbv2.SetIpAddressTypeInput, ...request.Option) (*elbv2.SetIpAddressTypeOutput, error)
	SetIpAddressTypeRequest(*elbv2.SetIpAddressTypeInput) (*request.Request, *elbv2.SetIpAddressTypeOutput)

	SetRulePriorities(*elbv2.SetRulePrioritiesInput) (*elbv2.SetRulePrioritiesOutput, error)
	SetRulePrioritiesWithContext(aws.Context, *elbv2.SetRulePrioritiesInput, ...request.Option) (*elbv2.SetRulePrioritiesOutput, error)
	SetRulePrioritiesRequest(*elbv2.SetRulePrioritiesInput) (*request.Request, *elbv2.SetRulePrioritiesOutput)

	SetSecurityGroups(*elbv2.SetSecurityGroupsInput) (*elbv2.SetSecurityGroupsOutput, error)
	SetSecurityGroupsWithContext(aws.Context, *elbv2.SetSecurityGroupsInput, ...request.Option) (*elbv2.SetSecurityGroupsOutput, error)
	SetSecurityGroupsRequest(*elbv2.SetSecurityGroupsInput) (*request.Request, *elbv2.SetSecurityGroupsOutput)

	SetSubnets(*elbv2.SetSubnetsInput) (*elbv2.SetSubnetsOutput, error)
	SetSubnetsWithContext(aws.Context, *elbv2.SetSubnetsInput, ...request.Option) (*elbv2.SetSubnetsOutput, error)
	SetSubnetsRequest(*elbv2.SetSubnetsInput) (*request.Request, *elbv2.SetSubnetsOutput)

	WaitUntilLoadBalancerAvailable(*elbv2.DescribeLoadBalancersInput) error
	WaitUntilLoadBalancerAvailableWithContext(aws.Context, *elbv2.DescribeLoadBalancersInput, ...request.WaiterOption) error

	WaitUntilLoadBalancerExists(*elbv2.DescribeLoadBalancersInput) error
	WaitUntilLoadBalancerExistsWithContext(aws.Context, *elbv2.DescribeLoadBalancersInput, ...request.WaiterOption) error

	WaitUntilLoadBalancersDeleted(*elbv2.DescribeLoadBalancersInput) error
	WaitUntilLoadBalancersDeletedWithContext(aws.Context, *elbv2.DescribeLoadBalancersInput, ...request.WaiterOption) error
}

var _ ELBV2API = (*elbv2.ELBV2)(nil)
