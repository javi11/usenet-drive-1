// Code generated by MockGen. DO NOT EDIT.
// Source: ./connection.go

// Package nntpcli is a generated GoMock package.
package nntpcli

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockConnection is a mock of Connection interface.
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection.
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance.
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockConnection) Authenticate(username, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", username, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockConnectionMockRecorder) Authenticate(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockConnection)(nil).Authenticate), username, password)
}

// Body mocks base method.
func (m *MockConnection) Body(id string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Body", id)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Body indicates an expected call of Body.
func (mr *MockConnectionMockRecorder) Body(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockConnection)(nil).Body), id)
}

// Post mocks base method.
func (m *MockConnection) Post(p []byte, chunkSize int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", p, chunkSize)
	ret0, _ := ret[0].(error)
	return ret0
}

// Post indicates an expected call of Post.
func (mr *MockConnectionMockRecorder) Post(p, chunkSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockConnection)(nil).Post), p, chunkSize)
}

// ProviderID mocks base method.
func (m *MockConnection) ProviderID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProviderID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ProviderID indicates an expected call of ProviderID.
func (mr *MockConnectionMockRecorder) ProviderID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProviderID", reflect.TypeOf((*MockConnection)(nil).ProviderID))
}

// Quit mocks base method.
func (m *MockConnection) Quit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Quit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Quit indicates an expected call of Quit.
func (mr *MockConnectionMockRecorder) Quit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quit", reflect.TypeOf((*MockConnection)(nil).Quit))
}

// SelectGroup mocks base method.
func (m *MockConnection) SelectGroup(group string) (int, int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectGroup", group)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(int)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// SelectGroup indicates an expected call of SelectGroup.
func (mr *MockConnectionMockRecorder) SelectGroup(group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectGroup", reflect.TypeOf((*MockConnection)(nil).SelectGroup), group)
}
