// Code generated by MockGen. DO NOT EDIT.
// Source: hashing.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHashing is a mock of Hashing interface.
type MockHashing struct {
	ctrl     *gomock.Controller
	recorder *MockHashingMockRecorder
}

// MockHashingMockRecorder is the mock recorder for MockHashing.
type MockHashingMockRecorder struct {
	mock *MockHashing
}

// NewMockHashing creates a new mock instance.
func NewMockHashing(ctrl *gomock.Controller) *MockHashing {
	mock := &MockHashing{ctrl: ctrl}
	mock.recorder = &MockHashingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHashing) EXPECT() *MockHashingMockRecorder {
	return m.recorder
}

// Hash mocks base method.
func (m *MockHashing) Hash(input string) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash", input)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Hash indicates an expected call of Hash.
func (mr *MockHashingMockRecorder) Hash(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockHashing)(nil).Hash), input)
}
