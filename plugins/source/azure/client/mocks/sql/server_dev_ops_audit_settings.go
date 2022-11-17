// Code generated by MockGen. DO NOT EDIT.
// Source: server_dev_ops_audit_settings.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	armsql "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	sql "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/sql"
	gomock "github.com/golang/mock/gomock"
)

// MockServerDevOpsAuditSettingsClient is a mock of ServerDevOpsAuditSettingsClient interface.
type MockServerDevOpsAuditSettingsClient struct {
	ctrl     *gomock.Controller
	recorder *MockServerDevOpsAuditSettingsClientMockRecorder
}

// MockServerDevOpsAuditSettingsClientMockRecorder is the mock recorder for MockServerDevOpsAuditSettingsClient.
type MockServerDevOpsAuditSettingsClientMockRecorder struct {
	mock *MockServerDevOpsAuditSettingsClient
}

// NewMockServerDevOpsAuditSettingsClient creates a new mock instance.
func NewMockServerDevOpsAuditSettingsClient(ctrl *gomock.Controller) *MockServerDevOpsAuditSettingsClient {
	mock := &MockServerDevOpsAuditSettingsClient{ctrl: ctrl}
	mock.recorder = &MockServerDevOpsAuditSettingsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerDevOpsAuditSettingsClient) EXPECT() *MockServerDevOpsAuditSettingsClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockServerDevOpsAuditSettingsClient) Get(arg0 context.Context, arg1, arg2, arg3 string, arg4 *armsql.ServerDevOpsAuditSettingsClientGetOptions) (armsql.ServerDevOpsAuditSettingsClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(armsql.ServerDevOpsAuditSettingsClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockServerDevOpsAuditSettingsClientMockRecorder) Get(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockServerDevOpsAuditSettingsClient)(nil).Get), arg0, arg1, arg2, arg3, arg4)
}

// NewListByServerPager mocks base method.
func (m *MockServerDevOpsAuditSettingsClient) NewListByServerPager(arg0, arg1 string, arg2 *armsql.ServerDevOpsAuditSettingsClientListByServerOptions) *sql.RuntimePagerArmsqlServerDevOpsAuditSettingsClientListByServerResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListByServerPager", arg0, arg1, arg2)
	ret0, _ := ret[0].(*sql.RuntimePagerArmsqlServerDevOpsAuditSettingsClientListByServerResponse)
	return ret0
}

// NewListByServerPager indicates an expected call of NewListByServerPager.
func (mr *MockServerDevOpsAuditSettingsClientMockRecorder) NewListByServerPager(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListByServerPager", reflect.TypeOf((*MockServerDevOpsAuditSettingsClient)(nil).NewListByServerPager), arg0, arg1, arg2)
}
