// Code generated by MockGen. DO NOT EDIT.
// Source: managed_instances.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	armsql "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	sql "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/sql"
	gomock "github.com/golang/mock/gomock"
)

// MockManagedInstancesClient is a mock of ManagedInstancesClient interface.
type MockManagedInstancesClient struct {
	ctrl     *gomock.Controller
	recorder *MockManagedInstancesClientMockRecorder
}

// MockManagedInstancesClientMockRecorder is the mock recorder for MockManagedInstancesClient.
type MockManagedInstancesClientMockRecorder struct {
	mock *MockManagedInstancesClient
}

// NewMockManagedInstancesClient creates a new mock instance.
func NewMockManagedInstancesClient(ctrl *gomock.Controller) *MockManagedInstancesClient {
	mock := &MockManagedInstancesClient{ctrl: ctrl}
	mock.recorder = &MockManagedInstancesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManagedInstancesClient) EXPECT() *MockManagedInstancesClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockManagedInstancesClient) Get(arg0 context.Context, arg1, arg2 string, arg3 *armsql.ManagedInstancesClientGetOptions) (armsql.ManagedInstancesClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(armsql.ManagedInstancesClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockManagedInstancesClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockManagedInstancesClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// NewListByInstancePoolPager mocks base method.
func (m *MockManagedInstancesClient) NewListByInstancePoolPager(arg0, arg1 string, arg2 *armsql.ManagedInstancesClientListByInstancePoolOptions) *sql.RuntimePagerArmsqlManagedInstancesClientListByInstancePoolResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListByInstancePoolPager", arg0, arg1, arg2)
	ret0, _ := ret[0].(*sql.RuntimePagerArmsqlManagedInstancesClientListByInstancePoolResponse)
	return ret0
}

// NewListByInstancePoolPager indicates an expected call of NewListByInstancePoolPager.
func (mr *MockManagedInstancesClientMockRecorder) NewListByInstancePoolPager(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListByInstancePoolPager", reflect.TypeOf((*MockManagedInstancesClient)(nil).NewListByInstancePoolPager), arg0, arg1, arg2)
}

// NewListByManagedInstancePager mocks base method.
func (m *MockManagedInstancesClient) NewListByManagedInstancePager(arg0, arg1 string, arg2 *armsql.ManagedInstancesClientListByManagedInstanceOptions) *sql.RuntimePagerArmsqlManagedInstancesClientListByManagedInstanceResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListByManagedInstancePager", arg0, arg1, arg2)
	ret0, _ := ret[0].(*sql.RuntimePagerArmsqlManagedInstancesClientListByManagedInstanceResponse)
	return ret0
}

// NewListByManagedInstancePager indicates an expected call of NewListByManagedInstancePager.
func (mr *MockManagedInstancesClientMockRecorder) NewListByManagedInstancePager(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListByManagedInstancePager", reflect.TypeOf((*MockManagedInstancesClient)(nil).NewListByManagedInstancePager), arg0, arg1, arg2)
}

// NewListByResourceGroupPager mocks base method.
func (m *MockManagedInstancesClient) NewListByResourceGroupPager(arg0 string, arg1 *armsql.ManagedInstancesClientListByResourceGroupOptions) *sql.RuntimePagerArmsqlManagedInstancesClientListByResourceGroupResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListByResourceGroupPager", arg0, arg1)
	ret0, _ := ret[0].(*sql.RuntimePagerArmsqlManagedInstancesClientListByResourceGroupResponse)
	return ret0
}

// NewListByResourceGroupPager indicates an expected call of NewListByResourceGroupPager.
func (mr *MockManagedInstancesClientMockRecorder) NewListByResourceGroupPager(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListByResourceGroupPager", reflect.TypeOf((*MockManagedInstancesClient)(nil).NewListByResourceGroupPager), arg0, arg1)
}

// NewListPager mocks base method.
func (m *MockManagedInstancesClient) NewListPager(arg0 *armsql.ManagedInstancesClientListOptions) *sql.RuntimePagerArmsqlManagedInstancesClientListResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", arg0)
	ret0, _ := ret[0].(*sql.RuntimePagerArmsqlManagedInstancesClientListResponse)
	return ret0
}

// NewListPager indicates an expected call of NewListPager.
func (mr *MockManagedInstancesClientMockRecorder) NewListPager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockManagedInstancesClient)(nil).NewListPager), arg0)
}
