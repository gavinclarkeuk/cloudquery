// Code generated by MockGen. DO NOT EDIT.
// Source: subscriptions.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	armsubscription "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	subscription "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/subscription"
	gomock "github.com/golang/mock/gomock"
)

// MockSubscriptionsClient is a mock of SubscriptionsClient interface.
type MockSubscriptionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionsClientMockRecorder
}

// MockSubscriptionsClientMockRecorder is the mock recorder for MockSubscriptionsClient.
type MockSubscriptionsClientMockRecorder struct {
	mock *MockSubscriptionsClient
}

// NewMockSubscriptionsClient creates a new mock instance.
func NewMockSubscriptionsClient(ctrl *gomock.Controller) *MockSubscriptionsClient {
	mock := &MockSubscriptionsClient{ctrl: ctrl}
	mock.recorder = &MockSubscriptionsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscriptionsClient) EXPECT() *MockSubscriptionsClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockSubscriptionsClient) Get(arg0 context.Context, arg1 string, arg2 *armsubscription.SubscriptionsClientGetOptions) (armsubscription.SubscriptionsClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(armsubscription.SubscriptionsClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSubscriptionsClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSubscriptionsClient)(nil).Get), arg0, arg1, arg2)
}

// NewListLocationsPager mocks base method.
func (m *MockSubscriptionsClient) NewListLocationsPager(arg0 string, arg1 *armsubscription.SubscriptionsClientListLocationsOptions) *subscription.RuntimePagerArmsubscriptionSubscriptionsClientListLocationsResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListLocationsPager", arg0, arg1)
	ret0, _ := ret[0].(*subscription.RuntimePagerArmsubscriptionSubscriptionsClientListLocationsResponse)
	return ret0
}

// NewListLocationsPager indicates an expected call of NewListLocationsPager.
func (mr *MockSubscriptionsClientMockRecorder) NewListLocationsPager(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListLocationsPager", reflect.TypeOf((*MockSubscriptionsClient)(nil).NewListLocationsPager), arg0, arg1)
}

// NewListPager mocks base method.
func (m *MockSubscriptionsClient) NewListPager(arg0 *armsubscription.SubscriptionsClientListOptions) *subscription.RuntimePagerArmsubscriptionSubscriptionsClientListResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", arg0)
	ret0, _ := ret[0].(*subscription.RuntimePagerArmsubscriptionSubscriptionsClientListResponse)
	return ret0
}

// NewListPager indicates an expected call of NewListPager.
func (mr *MockSubscriptionsClientMockRecorder) NewListPager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockSubscriptionsClient)(nil).NewListPager), arg0)
}
