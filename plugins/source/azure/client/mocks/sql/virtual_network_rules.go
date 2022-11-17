// Code generated by MockGen. DO NOT EDIT.
// Source: virtual_network_rules.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	armsql "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	sql "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/sql"
	gomock "github.com/golang/mock/gomock"
)

// MockVirtualNetworkRulesClient is a mock of VirtualNetworkRulesClient interface.
type MockVirtualNetworkRulesClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualNetworkRulesClientMockRecorder
}

// MockVirtualNetworkRulesClientMockRecorder is the mock recorder for MockVirtualNetworkRulesClient.
type MockVirtualNetworkRulesClientMockRecorder struct {
	mock *MockVirtualNetworkRulesClient
}

// NewMockVirtualNetworkRulesClient creates a new mock instance.
func NewMockVirtualNetworkRulesClient(ctrl *gomock.Controller) *MockVirtualNetworkRulesClient {
	mock := &MockVirtualNetworkRulesClient{ctrl: ctrl}
	mock.recorder = &MockVirtualNetworkRulesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualNetworkRulesClient) EXPECT() *MockVirtualNetworkRulesClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockVirtualNetworkRulesClient) Get(arg0 context.Context, arg1, arg2, arg3 string, arg4 *armsql.VirtualNetworkRulesClientGetOptions) (armsql.VirtualNetworkRulesClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(armsql.VirtualNetworkRulesClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockVirtualNetworkRulesClientMockRecorder) Get(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVirtualNetworkRulesClient)(nil).Get), arg0, arg1, arg2, arg3, arg4)
}

// NewListByServerPager mocks base method.
func (m *MockVirtualNetworkRulesClient) NewListByServerPager(arg0, arg1 string, arg2 *armsql.VirtualNetworkRulesClientListByServerOptions) *sql.RuntimePagerArmsqlVirtualNetworkRulesClientListByServerResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListByServerPager", arg0, arg1, arg2)
	ret0, _ := ret[0].(*sql.RuntimePagerArmsqlVirtualNetworkRulesClientListByServerResponse)
	return ret0
}

// NewListByServerPager indicates an expected call of NewListByServerPager.
func (mr *MockVirtualNetworkRulesClientMockRecorder) NewListByServerPager(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListByServerPager", reflect.TypeOf((*MockVirtualNetworkRulesClient)(nil).NewListByServerPager), arg0, arg1, arg2)
}
