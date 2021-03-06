// Code generated by MockGen. DO NOT EDIT.
// Source: ./membership.go

// Package mock_internal_watcher is a generated GoMock package.
package mock_internal_watcher

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockClusterSecretHandler is a mock of ClusterSecretHandler interface.
type MockClusterSecretHandler struct {
	ctrl     *gomock.Controller
	recorder *MockClusterSecretHandlerMockRecorder
}

// MockClusterSecretHandlerMockRecorder is the mock recorder for MockClusterSecretHandler.
type MockClusterSecretHandlerMockRecorder struct {
	mock *MockClusterSecretHandler
}

// NewMockClusterSecretHandler creates a new mock instance.
func NewMockClusterSecretHandler(ctrl *gomock.Controller) *MockClusterSecretHandler {
	mock := &MockClusterSecretHandler{ctrl: ctrl}
	mock.recorder = &MockClusterSecretHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterSecretHandler) EXPECT() *MockClusterSecretHandlerMockRecorder {
	return m.recorder
}

// AddMemberCluster mocks base method.
func (m *MockClusterSecretHandler) AddMemberCluster(ctx context.Context, s *v1.Secret) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMemberCluster", ctx, s)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMemberCluster indicates an expected call of AddMemberCluster.
func (mr *MockClusterSecretHandlerMockRecorder) AddMemberCluster(ctx, s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMemberCluster", reflect.TypeOf((*MockClusterSecretHandler)(nil).AddMemberCluster), ctx, s)
}

// DeleteMemberCluster mocks base method.
func (m *MockClusterSecretHandler) DeleteMemberCluster(ctx context.Context, s *v1.Secret) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMemberCluster", ctx, s)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMemberCluster indicates an expected call of DeleteMemberCluster.
func (mr *MockClusterSecretHandlerMockRecorder) DeleteMemberCluster(ctx, s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMemberCluster", reflect.TypeOf((*MockClusterSecretHandler)(nil).DeleteMemberCluster), ctx, s)
}
