// Code generated by MockGen. DO NOT EDIT.
// Source: wafv2.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	wafv2 "github.com/aws/aws-sdk-go-v2/service/wafv2"
	gomock "github.com/golang/mock/gomock"
)

// MockWafv2Client is a mock of Wafv2Client interface.
type MockWafv2Client struct {
	ctrl     *gomock.Controller
	recorder *MockWafv2ClientMockRecorder
}

// MockWafv2ClientMockRecorder is the mock recorder for MockWafv2Client.
type MockWafv2ClientMockRecorder struct {
	mock *MockWafv2Client
}

// NewMockWafv2Client creates a new mock instance.
func NewMockWafv2Client(ctrl *gomock.Controller) *MockWafv2Client {
	mock := &MockWafv2Client{ctrl: ctrl}
	mock.recorder = &MockWafv2ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWafv2Client) EXPECT() *MockWafv2ClientMockRecorder {
	return m.recorder
}

// DescribeManagedRuleGroup mocks base method.
func (m *MockWafv2Client) DescribeManagedRuleGroup(arg0 context.Context, arg1 *wafv2.DescribeManagedRuleGroupInput, arg2 ...func(*wafv2.Options)) (*wafv2.DescribeManagedRuleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeManagedRuleGroup", varargs...)
	ret0, _ := ret[0].(*wafv2.DescribeManagedRuleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeManagedRuleGroup indicates an expected call of DescribeManagedRuleGroup.
func (mr *MockWafv2ClientMockRecorder) DescribeManagedRuleGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeManagedRuleGroup", reflect.TypeOf((*MockWafv2Client)(nil).DescribeManagedRuleGroup), varargs...)
}

// GetIPSet mocks base method.
func (m *MockWafv2Client) GetIPSet(arg0 context.Context, arg1 *wafv2.GetIPSetInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetIPSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIPSet", varargs...)
	ret0, _ := ret[0].(*wafv2.GetIPSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIPSet indicates an expected call of GetIPSet.
func (mr *MockWafv2ClientMockRecorder) GetIPSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPSet", reflect.TypeOf((*MockWafv2Client)(nil).GetIPSet), varargs...)
}

// GetLoggingConfiguration mocks base method.
func (m *MockWafv2Client) GetLoggingConfiguration(arg0 context.Context, arg1 *wafv2.GetLoggingConfigurationInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetLoggingConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLoggingConfiguration", varargs...)
	ret0, _ := ret[0].(*wafv2.GetLoggingConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoggingConfiguration indicates an expected call of GetLoggingConfiguration.
func (mr *MockWafv2ClientMockRecorder) GetLoggingConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoggingConfiguration", reflect.TypeOf((*MockWafv2Client)(nil).GetLoggingConfiguration), varargs...)
}

// GetManagedRuleSet mocks base method.
func (m *MockWafv2Client) GetManagedRuleSet(arg0 context.Context, arg1 *wafv2.GetManagedRuleSetInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetManagedRuleSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetManagedRuleSet", varargs...)
	ret0, _ := ret[0].(*wafv2.GetManagedRuleSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManagedRuleSet indicates an expected call of GetManagedRuleSet.
func (mr *MockWafv2ClientMockRecorder) GetManagedRuleSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedRuleSet", reflect.TypeOf((*MockWafv2Client)(nil).GetManagedRuleSet), varargs...)
}

// GetMobileSdkRelease mocks base method.
func (m *MockWafv2Client) GetMobileSdkRelease(arg0 context.Context, arg1 *wafv2.GetMobileSdkReleaseInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetMobileSdkReleaseOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMobileSdkRelease", varargs...)
	ret0, _ := ret[0].(*wafv2.GetMobileSdkReleaseOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMobileSdkRelease indicates an expected call of GetMobileSdkRelease.
func (mr *MockWafv2ClientMockRecorder) GetMobileSdkRelease(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMobileSdkRelease", reflect.TypeOf((*MockWafv2Client)(nil).GetMobileSdkRelease), varargs...)
}

// GetPermissionPolicy mocks base method.
func (m *MockWafv2Client) GetPermissionPolicy(arg0 context.Context, arg1 *wafv2.GetPermissionPolicyInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetPermissionPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPermissionPolicy", varargs...)
	ret0, _ := ret[0].(*wafv2.GetPermissionPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPermissionPolicy indicates an expected call of GetPermissionPolicy.
func (mr *MockWafv2ClientMockRecorder) GetPermissionPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPermissionPolicy", reflect.TypeOf((*MockWafv2Client)(nil).GetPermissionPolicy), varargs...)
}

// GetRateBasedStatementManagedKeys mocks base method.
func (m *MockWafv2Client) GetRateBasedStatementManagedKeys(arg0 context.Context, arg1 *wafv2.GetRateBasedStatementManagedKeysInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetRateBasedStatementManagedKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRateBasedStatementManagedKeys", varargs...)
	ret0, _ := ret[0].(*wafv2.GetRateBasedStatementManagedKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRateBasedStatementManagedKeys indicates an expected call of GetRateBasedStatementManagedKeys.
func (mr *MockWafv2ClientMockRecorder) GetRateBasedStatementManagedKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRateBasedStatementManagedKeys", reflect.TypeOf((*MockWafv2Client)(nil).GetRateBasedStatementManagedKeys), varargs...)
}

// GetRegexPatternSet mocks base method.
func (m *MockWafv2Client) GetRegexPatternSet(arg0 context.Context, arg1 *wafv2.GetRegexPatternSetInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetRegexPatternSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRegexPatternSet", varargs...)
	ret0, _ := ret[0].(*wafv2.GetRegexPatternSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegexPatternSet indicates an expected call of GetRegexPatternSet.
func (mr *MockWafv2ClientMockRecorder) GetRegexPatternSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegexPatternSet", reflect.TypeOf((*MockWafv2Client)(nil).GetRegexPatternSet), varargs...)
}

// GetRuleGroup mocks base method.
func (m *MockWafv2Client) GetRuleGroup(arg0 context.Context, arg1 *wafv2.GetRuleGroupInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetRuleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRuleGroup", varargs...)
	ret0, _ := ret[0].(*wafv2.GetRuleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRuleGroup indicates an expected call of GetRuleGroup.
func (mr *MockWafv2ClientMockRecorder) GetRuleGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRuleGroup", reflect.TypeOf((*MockWafv2Client)(nil).GetRuleGroup), varargs...)
}

// GetSampledRequests mocks base method.
func (m *MockWafv2Client) GetSampledRequests(arg0 context.Context, arg1 *wafv2.GetSampledRequestsInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetSampledRequestsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSampledRequests", varargs...)
	ret0, _ := ret[0].(*wafv2.GetSampledRequestsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSampledRequests indicates an expected call of GetSampledRequests.
func (mr *MockWafv2ClientMockRecorder) GetSampledRequests(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSampledRequests", reflect.TypeOf((*MockWafv2Client)(nil).GetSampledRequests), varargs...)
}

// GetWebACL mocks base method.
func (m *MockWafv2Client) GetWebACL(arg0 context.Context, arg1 *wafv2.GetWebACLInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetWebACLOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWebACL", varargs...)
	ret0, _ := ret[0].(*wafv2.GetWebACLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWebACL indicates an expected call of GetWebACL.
func (mr *MockWafv2ClientMockRecorder) GetWebACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWebACL", reflect.TypeOf((*MockWafv2Client)(nil).GetWebACL), varargs...)
}

// GetWebACLForResource mocks base method.
func (m *MockWafv2Client) GetWebACLForResource(arg0 context.Context, arg1 *wafv2.GetWebACLForResourceInput, arg2 ...func(*wafv2.Options)) (*wafv2.GetWebACLForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWebACLForResource", varargs...)
	ret0, _ := ret[0].(*wafv2.GetWebACLForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWebACLForResource indicates an expected call of GetWebACLForResource.
func (mr *MockWafv2ClientMockRecorder) GetWebACLForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWebACLForResource", reflect.TypeOf((*MockWafv2Client)(nil).GetWebACLForResource), varargs...)
}

// ListAvailableManagedRuleGroupVersions mocks base method.
func (m *MockWafv2Client) ListAvailableManagedRuleGroupVersions(arg0 context.Context, arg1 *wafv2.ListAvailableManagedRuleGroupVersionsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListAvailableManagedRuleGroupVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAvailableManagedRuleGroupVersions", varargs...)
	ret0, _ := ret[0].(*wafv2.ListAvailableManagedRuleGroupVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAvailableManagedRuleGroupVersions indicates an expected call of ListAvailableManagedRuleGroupVersions.
func (mr *MockWafv2ClientMockRecorder) ListAvailableManagedRuleGroupVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAvailableManagedRuleGroupVersions", reflect.TypeOf((*MockWafv2Client)(nil).ListAvailableManagedRuleGroupVersions), varargs...)
}

// ListAvailableManagedRuleGroups mocks base method.
func (m *MockWafv2Client) ListAvailableManagedRuleGroups(arg0 context.Context, arg1 *wafv2.ListAvailableManagedRuleGroupsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListAvailableManagedRuleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAvailableManagedRuleGroups", varargs...)
	ret0, _ := ret[0].(*wafv2.ListAvailableManagedRuleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAvailableManagedRuleGroups indicates an expected call of ListAvailableManagedRuleGroups.
func (mr *MockWafv2ClientMockRecorder) ListAvailableManagedRuleGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAvailableManagedRuleGroups", reflect.TypeOf((*MockWafv2Client)(nil).ListAvailableManagedRuleGroups), varargs...)
}

// ListIPSets mocks base method.
func (m *MockWafv2Client) ListIPSets(arg0 context.Context, arg1 *wafv2.ListIPSetsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListIPSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIPSets", varargs...)
	ret0, _ := ret[0].(*wafv2.ListIPSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIPSets indicates an expected call of ListIPSets.
func (mr *MockWafv2ClientMockRecorder) ListIPSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIPSets", reflect.TypeOf((*MockWafv2Client)(nil).ListIPSets), varargs...)
}

// ListLoggingConfigurations mocks base method.
func (m *MockWafv2Client) ListLoggingConfigurations(arg0 context.Context, arg1 *wafv2.ListLoggingConfigurationsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListLoggingConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListLoggingConfigurations", varargs...)
	ret0, _ := ret[0].(*wafv2.ListLoggingConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLoggingConfigurations indicates an expected call of ListLoggingConfigurations.
func (mr *MockWafv2ClientMockRecorder) ListLoggingConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLoggingConfigurations", reflect.TypeOf((*MockWafv2Client)(nil).ListLoggingConfigurations), varargs...)
}

// ListManagedRuleSets mocks base method.
func (m *MockWafv2Client) ListManagedRuleSets(arg0 context.Context, arg1 *wafv2.ListManagedRuleSetsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListManagedRuleSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListManagedRuleSets", varargs...)
	ret0, _ := ret[0].(*wafv2.ListManagedRuleSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListManagedRuleSets indicates an expected call of ListManagedRuleSets.
func (mr *MockWafv2ClientMockRecorder) ListManagedRuleSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListManagedRuleSets", reflect.TypeOf((*MockWafv2Client)(nil).ListManagedRuleSets), varargs...)
}

// ListMobileSdkReleases mocks base method.
func (m *MockWafv2Client) ListMobileSdkReleases(arg0 context.Context, arg1 *wafv2.ListMobileSdkReleasesInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListMobileSdkReleasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMobileSdkReleases", varargs...)
	ret0, _ := ret[0].(*wafv2.ListMobileSdkReleasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMobileSdkReleases indicates an expected call of ListMobileSdkReleases.
func (mr *MockWafv2ClientMockRecorder) ListMobileSdkReleases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMobileSdkReleases", reflect.TypeOf((*MockWafv2Client)(nil).ListMobileSdkReleases), varargs...)
}

// ListRegexPatternSets mocks base method.
func (m *MockWafv2Client) ListRegexPatternSets(arg0 context.Context, arg1 *wafv2.ListRegexPatternSetsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListRegexPatternSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRegexPatternSets", varargs...)
	ret0, _ := ret[0].(*wafv2.ListRegexPatternSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRegexPatternSets indicates an expected call of ListRegexPatternSets.
func (mr *MockWafv2ClientMockRecorder) ListRegexPatternSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRegexPatternSets", reflect.TypeOf((*MockWafv2Client)(nil).ListRegexPatternSets), varargs...)
}

// ListResourcesForWebACL mocks base method.
func (m *MockWafv2Client) ListResourcesForWebACL(arg0 context.Context, arg1 *wafv2.ListResourcesForWebACLInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListResourcesForWebACLOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResourcesForWebACL", varargs...)
	ret0, _ := ret[0].(*wafv2.ListResourcesForWebACLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourcesForWebACL indicates an expected call of ListResourcesForWebACL.
func (mr *MockWafv2ClientMockRecorder) ListResourcesForWebACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourcesForWebACL", reflect.TypeOf((*MockWafv2Client)(nil).ListResourcesForWebACL), varargs...)
}

// ListRuleGroups mocks base method.
func (m *MockWafv2Client) ListRuleGroups(arg0 context.Context, arg1 *wafv2.ListRuleGroupsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListRuleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRuleGroups", varargs...)
	ret0, _ := ret[0].(*wafv2.ListRuleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRuleGroups indicates an expected call of ListRuleGroups.
func (mr *MockWafv2ClientMockRecorder) ListRuleGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRuleGroups", reflect.TypeOf((*MockWafv2Client)(nil).ListRuleGroups), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockWafv2Client) ListTagsForResource(arg0 context.Context, arg1 *wafv2.ListTagsForResourceInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*wafv2.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockWafv2ClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockWafv2Client)(nil).ListTagsForResource), varargs...)
}

// ListWebACLs mocks base method.
func (m *MockWafv2Client) ListWebACLs(arg0 context.Context, arg1 *wafv2.ListWebACLsInput, arg2 ...func(*wafv2.Options)) (*wafv2.ListWebACLsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWebACLs", varargs...)
	ret0, _ := ret[0].(*wafv2.ListWebACLsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWebACLs indicates an expected call of ListWebACLs.
func (mr *MockWafv2ClientMockRecorder) ListWebACLs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWebACLs", reflect.TypeOf((*MockWafv2Client)(nil).ListWebACLs), varargs...)
}