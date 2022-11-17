// Code generated by MockGen. DO NOT EDIT.
// Source: virtual_machine_scale_sets.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	armcompute "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	compute "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/compute"
	gomock "github.com/golang/mock/gomock"
)

// MockVirtualMachineScaleSetsClient is a mock of VirtualMachineScaleSetsClient interface.
type MockVirtualMachineScaleSetsClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMachineScaleSetsClientMockRecorder
}

// MockVirtualMachineScaleSetsClientMockRecorder is the mock recorder for MockVirtualMachineScaleSetsClient.
type MockVirtualMachineScaleSetsClientMockRecorder struct {
	mock *MockVirtualMachineScaleSetsClient
}

// NewMockVirtualMachineScaleSetsClient creates a new mock instance.
func NewMockVirtualMachineScaleSetsClient(ctrl *gomock.Controller) *MockVirtualMachineScaleSetsClient {
	mock := &MockVirtualMachineScaleSetsClient{ctrl: ctrl}
	mock.recorder = &MockVirtualMachineScaleSetsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMachineScaleSetsClient) EXPECT() *MockVirtualMachineScaleSetsClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockVirtualMachineScaleSetsClient) Get(arg0 context.Context, arg1, arg2 string, arg3 *armcompute.VirtualMachineScaleSetsClientGetOptions) (armcompute.VirtualMachineScaleSetsClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(armcompute.VirtualMachineScaleSetsClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// GetInstanceView mocks base method.
func (m *MockVirtualMachineScaleSetsClient) GetInstanceView(arg0 context.Context, arg1, arg2 string, arg3 *armcompute.VirtualMachineScaleSetsClientGetInstanceViewOptions) (armcompute.VirtualMachineScaleSetsClientGetInstanceViewResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceView", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(armcompute.VirtualMachineScaleSetsClientGetInstanceViewResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstanceView indicates an expected call of GetInstanceView.
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) GetInstanceView(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceView", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).GetInstanceView), arg0, arg1, arg2, arg3)
}

// NewListAllPager mocks base method.
func (m *MockVirtualMachineScaleSetsClient) NewListAllPager(arg0 *armcompute.VirtualMachineScaleSetsClientListAllOptions) *compute.RuntimePagerArmcomputeVirtualMachineScaleSetsClientListAllResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListAllPager", arg0)
	ret0, _ := ret[0].(*compute.RuntimePagerArmcomputeVirtualMachineScaleSetsClientListAllResponse)
	return ret0
}

// NewListAllPager indicates an expected call of NewListAllPager.
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) NewListAllPager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListAllPager", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).NewListAllPager), arg0)
}

// NewListByLocationPager mocks base method.
func (m *MockVirtualMachineScaleSetsClient) NewListByLocationPager(arg0 string, arg1 *armcompute.VirtualMachineScaleSetsClientListByLocationOptions) *compute.RuntimePagerArmcomputeVirtualMachineScaleSetsClientListByLocationResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListByLocationPager", arg0, arg1)
	ret0, _ := ret[0].(*compute.RuntimePagerArmcomputeVirtualMachineScaleSetsClientListByLocationResponse)
	return ret0
}

// NewListByLocationPager indicates an expected call of NewListByLocationPager.
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) NewListByLocationPager(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListByLocationPager", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).NewListByLocationPager), arg0, arg1)
}

// NewListPager mocks base method.
func (m *MockVirtualMachineScaleSetsClient) NewListPager(arg0 string, arg1 *armcompute.VirtualMachineScaleSetsClientListOptions) *compute.RuntimePagerArmcomputeVirtualMachineScaleSetsClientListResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", arg0, arg1)
	ret0, _ := ret[0].(*compute.RuntimePagerArmcomputeVirtualMachineScaleSetsClientListResponse)
	return ret0
}

// NewListPager indicates an expected call of NewListPager.
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) NewListPager(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).NewListPager), arg0, arg1)
}

// NewListSKUsPager mocks base method.
func (m *MockVirtualMachineScaleSetsClient) NewListSKUsPager(arg0, arg1 string, arg2 *armcompute.VirtualMachineScaleSetsClientListSKUsOptions) *compute.RuntimePagerArmcomputeVirtualMachineScaleSetsClientListSKUsResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListSKUsPager", arg0, arg1, arg2)
	ret0, _ := ret[0].(*compute.RuntimePagerArmcomputeVirtualMachineScaleSetsClientListSKUsResponse)
	return ret0
}

// NewListSKUsPager indicates an expected call of NewListSKUsPager.
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) NewListSKUsPager(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListSKUsPager", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).NewListSKUsPager), arg0, arg1, arg2)
}
