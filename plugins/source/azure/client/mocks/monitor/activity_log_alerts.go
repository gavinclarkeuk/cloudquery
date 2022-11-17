// Code generated by MockGen. DO NOT EDIT.
// Source: activity_log_alerts.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	armmonitor "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	monitor "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/monitor"
	gomock "github.com/golang/mock/gomock"
)

// MockActivityLogAlertsClient is a mock of ActivityLogAlertsClient interface.
type MockActivityLogAlertsClient struct {
	ctrl     *gomock.Controller
	recorder *MockActivityLogAlertsClientMockRecorder
}

// MockActivityLogAlertsClientMockRecorder is the mock recorder for MockActivityLogAlertsClient.
type MockActivityLogAlertsClientMockRecorder struct {
	mock *MockActivityLogAlertsClient
}

// NewMockActivityLogAlertsClient creates a new mock instance.
func NewMockActivityLogAlertsClient(ctrl *gomock.Controller) *MockActivityLogAlertsClient {
	mock := &MockActivityLogAlertsClient{ctrl: ctrl}
	mock.recorder = &MockActivityLogAlertsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActivityLogAlertsClient) EXPECT() *MockActivityLogAlertsClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockActivityLogAlertsClient) Get(arg0 context.Context, arg1, arg2 string, arg3 *armmonitor.ActivityLogAlertsClientGetOptions) (armmonitor.ActivityLogAlertsClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(armmonitor.ActivityLogAlertsClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockActivityLogAlertsClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockActivityLogAlertsClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// NewListByResourceGroupPager mocks base method.
func (m *MockActivityLogAlertsClient) NewListByResourceGroupPager(arg0 string, arg1 *armmonitor.ActivityLogAlertsClientListByResourceGroupOptions) *monitor.RuntimePagerArmmonitorActivityLogAlertsClientListByResourceGroupResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListByResourceGroupPager", arg0, arg1)
	ret0, _ := ret[0].(*monitor.RuntimePagerArmmonitorActivityLogAlertsClientListByResourceGroupResponse)
	return ret0
}

// NewListByResourceGroupPager indicates an expected call of NewListByResourceGroupPager.
func (mr *MockActivityLogAlertsClientMockRecorder) NewListByResourceGroupPager(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListByResourceGroupPager", reflect.TypeOf((*MockActivityLogAlertsClient)(nil).NewListByResourceGroupPager), arg0, arg1)
}

// NewListBySubscriptionIDPager mocks base method.
func (m *MockActivityLogAlertsClient) NewListBySubscriptionIDPager(arg0 *armmonitor.ActivityLogAlertsClientListBySubscriptionIDOptions) *monitor.RuntimePagerArmmonitorActivityLogAlertsClientListBySubscriptionIDResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListBySubscriptionIDPager", arg0)
	ret0, _ := ret[0].(*monitor.RuntimePagerArmmonitorActivityLogAlertsClientListBySubscriptionIDResponse)
	return ret0
}

// NewListBySubscriptionIDPager indicates an expected call of NewListBySubscriptionIDPager.
func (mr *MockActivityLogAlertsClientMockRecorder) NewListBySubscriptionIDPager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListBySubscriptionIDPager", reflect.TypeOf((*MockActivityLogAlertsClient)(nil).NewListBySubscriptionIDPager), arg0)
}
