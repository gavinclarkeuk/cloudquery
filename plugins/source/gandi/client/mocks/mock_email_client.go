// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/gandi/client (interfaces: EmailClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
)

// MockEmailClient is a mock of EmailClient interface.
type MockEmailClient struct {
	ctrl     *gomock.Controller
	recorder *MockEmailClientMockRecorder
}

// MockEmailClientMockRecorder is the mock recorder for MockEmailClient.
type MockEmailClientMockRecorder struct {
	mock *MockEmailClient
}

// NewMockEmailClient creates a new mock instance.
func NewMockEmailClient(ctrl *gomock.Controller) *MockEmailClient {
	mock := &MockEmailClient{ctrl: ctrl}
	mock.recorder = &MockEmailClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailClient) EXPECT() *MockEmailClientMockRecorder {
	return m.recorder
}
