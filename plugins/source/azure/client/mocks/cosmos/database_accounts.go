// Code generated by MockGen. DO NOT EDIT.
// Source: database_accounts.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	armcosmos "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	cosmos "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/cosmos"
	gomock "github.com/golang/mock/gomock"
)

// MockDatabaseAccountsClient is a mock of DatabaseAccountsClient interface.
type MockDatabaseAccountsClient struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseAccountsClientMockRecorder
}

// MockDatabaseAccountsClientMockRecorder is the mock recorder for MockDatabaseAccountsClient.
type MockDatabaseAccountsClientMockRecorder struct {
	mock *MockDatabaseAccountsClient
}

// NewMockDatabaseAccountsClient creates a new mock instance.
func NewMockDatabaseAccountsClient(ctrl *gomock.Controller) *MockDatabaseAccountsClient {
	mock := &MockDatabaseAccountsClient{ctrl: ctrl}
	mock.recorder = &MockDatabaseAccountsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseAccountsClient) EXPECT() *MockDatabaseAccountsClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockDatabaseAccountsClient) Get(arg0 context.Context, arg1, arg2 string, arg3 *armcosmos.DatabaseAccountsClientGetOptions) (armcosmos.DatabaseAccountsClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(armcosmos.DatabaseAccountsClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDatabaseAccountsClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDatabaseAccountsClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// GetReadOnlyKeys mocks base method.
func (m *MockDatabaseAccountsClient) GetReadOnlyKeys(arg0 context.Context, arg1, arg2 string, arg3 *armcosmos.DatabaseAccountsClientGetReadOnlyKeysOptions) (armcosmos.DatabaseAccountsClientGetReadOnlyKeysResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReadOnlyKeys", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(armcosmos.DatabaseAccountsClientGetReadOnlyKeysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReadOnlyKeys indicates an expected call of GetReadOnlyKeys.
func (mr *MockDatabaseAccountsClientMockRecorder) GetReadOnlyKeys(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReadOnlyKeys", reflect.TypeOf((*MockDatabaseAccountsClient)(nil).GetReadOnlyKeys), arg0, arg1, arg2, arg3)
}

// ListConnectionStrings mocks base method.
func (m *MockDatabaseAccountsClient) ListConnectionStrings(arg0 context.Context, arg1, arg2 string, arg3 *armcosmos.DatabaseAccountsClientListConnectionStringsOptions) (armcosmos.DatabaseAccountsClientListConnectionStringsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListConnectionStrings", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(armcosmos.DatabaseAccountsClientListConnectionStringsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConnectionStrings indicates an expected call of ListConnectionStrings.
func (mr *MockDatabaseAccountsClientMockRecorder) ListConnectionStrings(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConnectionStrings", reflect.TypeOf((*MockDatabaseAccountsClient)(nil).ListConnectionStrings), arg0, arg1, arg2, arg3)
}

// ListKeys mocks base method.
func (m *MockDatabaseAccountsClient) ListKeys(arg0 context.Context, arg1, arg2 string, arg3 *armcosmos.DatabaseAccountsClientListKeysOptions) (armcosmos.DatabaseAccountsClientListKeysResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(armcosmos.DatabaseAccountsClientListKeysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKeys indicates an expected call of ListKeys.
func (mr *MockDatabaseAccountsClientMockRecorder) ListKeys(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockDatabaseAccountsClient)(nil).ListKeys), arg0, arg1, arg2, arg3)
}

// ListReadOnlyKeys mocks base method.
func (m *MockDatabaseAccountsClient) ListReadOnlyKeys(arg0 context.Context, arg1, arg2 string, arg3 *armcosmos.DatabaseAccountsClientListReadOnlyKeysOptions) (armcosmos.DatabaseAccountsClientListReadOnlyKeysResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReadOnlyKeys", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(armcosmos.DatabaseAccountsClientListReadOnlyKeysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReadOnlyKeys indicates an expected call of ListReadOnlyKeys.
func (mr *MockDatabaseAccountsClientMockRecorder) ListReadOnlyKeys(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReadOnlyKeys", reflect.TypeOf((*MockDatabaseAccountsClient)(nil).ListReadOnlyKeys), arg0, arg1, arg2, arg3)
}

// NewListByResourceGroupPager mocks base method.
func (m *MockDatabaseAccountsClient) NewListByResourceGroupPager(arg0 string, arg1 *armcosmos.DatabaseAccountsClientListByResourceGroupOptions) *cosmos.RuntimePagerArmcosmosDatabaseAccountsClientListByResourceGroupResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListByResourceGroupPager", arg0, arg1)
	ret0, _ := ret[0].(*cosmos.RuntimePagerArmcosmosDatabaseAccountsClientListByResourceGroupResponse)
	return ret0
}

// NewListByResourceGroupPager indicates an expected call of NewListByResourceGroupPager.
func (mr *MockDatabaseAccountsClientMockRecorder) NewListByResourceGroupPager(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListByResourceGroupPager", reflect.TypeOf((*MockDatabaseAccountsClient)(nil).NewListByResourceGroupPager), arg0, arg1)
}

// NewListMetricDefinitionsPager mocks base method.
func (m *MockDatabaseAccountsClient) NewListMetricDefinitionsPager(arg0, arg1 string, arg2 *armcosmos.DatabaseAccountsClientListMetricDefinitionsOptions) *cosmos.RuntimePagerArmcosmosDatabaseAccountsClientListMetricDefinitionsResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListMetricDefinitionsPager", arg0, arg1, arg2)
	ret0, _ := ret[0].(*cosmos.RuntimePagerArmcosmosDatabaseAccountsClientListMetricDefinitionsResponse)
	return ret0
}

// NewListMetricDefinitionsPager indicates an expected call of NewListMetricDefinitionsPager.
func (mr *MockDatabaseAccountsClientMockRecorder) NewListMetricDefinitionsPager(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListMetricDefinitionsPager", reflect.TypeOf((*MockDatabaseAccountsClient)(nil).NewListMetricDefinitionsPager), arg0, arg1, arg2)
}

// NewListMetricsPager mocks base method.
func (m *MockDatabaseAccountsClient) NewListMetricsPager(arg0, arg1, arg2 string, arg3 *armcosmos.DatabaseAccountsClientListMetricsOptions) *cosmos.RuntimePagerArmcosmosDatabaseAccountsClientListMetricsResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListMetricsPager", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*cosmos.RuntimePagerArmcosmosDatabaseAccountsClientListMetricsResponse)
	return ret0
}

// NewListMetricsPager indicates an expected call of NewListMetricsPager.
func (mr *MockDatabaseAccountsClientMockRecorder) NewListMetricsPager(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListMetricsPager", reflect.TypeOf((*MockDatabaseAccountsClient)(nil).NewListMetricsPager), arg0, arg1, arg2, arg3)
}

// NewListPager mocks base method.
func (m *MockDatabaseAccountsClient) NewListPager(arg0 *armcosmos.DatabaseAccountsClientListOptions) *cosmos.RuntimePagerArmcosmosDatabaseAccountsClientListResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", arg0)
	ret0, _ := ret[0].(*cosmos.RuntimePagerArmcosmosDatabaseAccountsClientListResponse)
	return ret0
}

// NewListPager indicates an expected call of NewListPager.
func (mr *MockDatabaseAccountsClientMockRecorder) NewListPager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockDatabaseAccountsClient)(nil).NewListPager), arg0)
}

// NewListUsagesPager mocks base method.
func (m *MockDatabaseAccountsClient) NewListUsagesPager(arg0, arg1 string, arg2 *armcosmos.DatabaseAccountsClientListUsagesOptions) *cosmos.RuntimePagerArmcosmosDatabaseAccountsClientListUsagesResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListUsagesPager", arg0, arg1, arg2)
	ret0, _ := ret[0].(*cosmos.RuntimePagerArmcosmosDatabaseAccountsClientListUsagesResponse)
	return ret0
}

// NewListUsagesPager indicates an expected call of NewListUsagesPager.
func (mr *MockDatabaseAccountsClientMockRecorder) NewListUsagesPager(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListUsagesPager", reflect.TypeOf((*MockDatabaseAccountsClient)(nil).NewListUsagesPager), arg0, arg1, arg2)
}
