// Code generated by MockGen. DO NOT EDIT.
// Source: firewall_rules.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	armpostgresql "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
	postgresql "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/postgresql"
	gomock "github.com/golang/mock/gomock"
)

// MockFirewallRulesClient is a mock of FirewallRulesClient interface.
type MockFirewallRulesClient struct {
	ctrl     *gomock.Controller
	recorder *MockFirewallRulesClientMockRecorder
}

// MockFirewallRulesClientMockRecorder is the mock recorder for MockFirewallRulesClient.
type MockFirewallRulesClientMockRecorder struct {
	mock *MockFirewallRulesClient
}

// NewMockFirewallRulesClient creates a new mock instance.
func NewMockFirewallRulesClient(ctrl *gomock.Controller) *MockFirewallRulesClient {
	mock := &MockFirewallRulesClient{ctrl: ctrl}
	mock.recorder = &MockFirewallRulesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFirewallRulesClient) EXPECT() *MockFirewallRulesClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockFirewallRulesClient) Get(arg0 context.Context, arg1, arg2, arg3 string, arg4 *armpostgresql.FirewallRulesClientGetOptions) (armpostgresql.FirewallRulesClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(armpostgresql.FirewallRulesClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockFirewallRulesClientMockRecorder) Get(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFirewallRulesClient)(nil).Get), arg0, arg1, arg2, arg3, arg4)
}

// NewListByServerPager mocks base method.
func (m *MockFirewallRulesClient) NewListByServerPager(arg0, arg1 string, arg2 *armpostgresql.FirewallRulesClientListByServerOptions) *postgresql.RuntimePagerArmpostgresqlFirewallRulesClientListByServerResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListByServerPager", arg0, arg1, arg2)
	ret0, _ := ret[0].(*postgresql.RuntimePagerArmpostgresqlFirewallRulesClientListByServerResponse)
	return ret0
}

// NewListByServerPager indicates an expected call of NewListByServerPager.
func (mr *MockFirewallRulesClientMockRecorder) NewListByServerPager(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListByServerPager", reflect.TypeOf((*MockFirewallRulesClient)(nil).NewListByServerPager), arg0, arg1, arg2)
}
