// Code generated by MockGen. DO NOT EDIT.
// Source: rules.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	armcdn "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	cdn "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/cdn"
	gomock "github.com/golang/mock/gomock"
)

// MockRulesClient is a mock of RulesClient interface.
type MockRulesClient struct {
	ctrl     *gomock.Controller
	recorder *MockRulesClientMockRecorder
}

// MockRulesClientMockRecorder is the mock recorder for MockRulesClient.
type MockRulesClientMockRecorder struct {
	mock *MockRulesClient
}

// NewMockRulesClient creates a new mock instance.
func NewMockRulesClient(ctrl *gomock.Controller) *MockRulesClient {
	mock := &MockRulesClient{ctrl: ctrl}
	mock.recorder = &MockRulesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRulesClient) EXPECT() *MockRulesClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockRulesClient) Get(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 *armcdn.RulesClientGetOptions) (armcdn.RulesClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(armcdn.RulesClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRulesClientMockRecorder) Get(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRulesClient)(nil).Get), arg0, arg1, arg2, arg3, arg4, arg5)
}

// NewListByRuleSetPager mocks base method.
func (m *MockRulesClient) NewListByRuleSetPager(arg0, arg1, arg2 string, arg3 *armcdn.RulesClientListByRuleSetOptions) *cdn.RuntimePagerArmcdnRulesClientListByRuleSetResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListByRuleSetPager", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*cdn.RuntimePagerArmcdnRulesClientListByRuleSetResponse)
	return ret0
}

// NewListByRuleSetPager indicates an expected call of NewListByRuleSetPager.
func (mr *MockRulesClientMockRecorder) NewListByRuleSetPager(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListByRuleSetPager", reflect.TypeOf((*MockRulesClient)(nil).NewListByRuleSetPager), arg0, arg1, arg2, arg3)
}
