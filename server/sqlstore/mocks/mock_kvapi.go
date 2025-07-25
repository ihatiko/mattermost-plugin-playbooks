// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-tech-develop/server/sqlstore (interfaces: KVAPI)

// Package mock_sqlstore is a generated GoMock package.
package mock_sqlstore

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockKVAPI is a mock of KVAPI interface.
type MockKVAPI struct {
	ctrl     *gomock.Controller
	recorder *MockKVAPIMockRecorder
}

// MockKVAPIMockRecorder is the mock recorder for MockKVAPI.
type MockKVAPIMockRecorder struct {
	mock *MockKVAPI
}

// NewMockKVAPI creates a new mock instance.
func NewMockKVAPI(ctrl *gomock.Controller) *MockKVAPI {
	mock := &MockKVAPI{ctrl: ctrl}
	mock.recorder = &MockKVAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKVAPI) EXPECT() *MockKVAPIMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockKVAPI) Get(arg0 string, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockKVAPIMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKVAPI)(nil).Get), arg0, arg1)
}
