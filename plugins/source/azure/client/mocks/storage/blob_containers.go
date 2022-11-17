// Code generated by MockGen. DO NOT EDIT.
// Source: blob_containers.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	armstorage "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	storage "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/storage"
	gomock "github.com/golang/mock/gomock"
)

// MockBlobContainersClient is a mock of BlobContainersClient interface.
type MockBlobContainersClient struct {
	ctrl     *gomock.Controller
	recorder *MockBlobContainersClientMockRecorder
}

// MockBlobContainersClientMockRecorder is the mock recorder for MockBlobContainersClient.
type MockBlobContainersClientMockRecorder struct {
	mock *MockBlobContainersClient
}

// NewMockBlobContainersClient creates a new mock instance.
func NewMockBlobContainersClient(ctrl *gomock.Controller) *MockBlobContainersClient {
	mock := &MockBlobContainersClient{ctrl: ctrl}
	mock.recorder = &MockBlobContainersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlobContainersClient) EXPECT() *MockBlobContainersClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockBlobContainersClient) Get(arg0 context.Context, arg1, arg2, arg3 string, arg4 *armstorage.BlobContainersClientGetOptions) (armstorage.BlobContainersClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(armstorage.BlobContainersClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBlobContainersClientMockRecorder) Get(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBlobContainersClient)(nil).Get), arg0, arg1, arg2, arg3, arg4)
}

// GetImmutabilityPolicy mocks base method.
func (m *MockBlobContainersClient) GetImmutabilityPolicy(arg0 context.Context, arg1, arg2, arg3 string, arg4 *armstorage.BlobContainersClientGetImmutabilityPolicyOptions) (armstorage.BlobContainersClientGetImmutabilityPolicyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImmutabilityPolicy", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(armstorage.BlobContainersClientGetImmutabilityPolicyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImmutabilityPolicy indicates an expected call of GetImmutabilityPolicy.
func (mr *MockBlobContainersClientMockRecorder) GetImmutabilityPolicy(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImmutabilityPolicy", reflect.TypeOf((*MockBlobContainersClient)(nil).GetImmutabilityPolicy), arg0, arg1, arg2, arg3, arg4)
}

// NewListPager mocks base method.
func (m *MockBlobContainersClient) NewListPager(arg0, arg1 string, arg2 *armstorage.BlobContainersClientListOptions) *storage.RuntimePagerArmstorageBlobContainersClientListResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.RuntimePagerArmstorageBlobContainersClientListResponse)
	return ret0
}

// NewListPager indicates an expected call of NewListPager.
func (mr *MockBlobContainersClientMockRecorder) NewListPager(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockBlobContainersClient)(nil).NewListPager), arg0, arg1, arg2)
}
