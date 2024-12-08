// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	partnerverifier "github.com/education-english-web/BE-english-web/pkg/partnerverifier"
)

// MockPartnerVerifier is a mock of PartnerVerifier interface.
type MockPartnerVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockPartnerVerifierMockRecorder
}

// MockPartnerVerifierMockRecorder is the mock recorder for MockPartnerVerifier.
type MockPartnerVerifierMockRecorder struct {
	mock *MockPartnerVerifier
}

// NewMockPartnerVerifier creates a new mock instance.
func NewMockPartnerVerifier(ctrl *gomock.Controller) *MockPartnerVerifier {
	mock := &MockPartnerVerifier{ctrl: ctrl}
	mock.recorder = &MockPartnerVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPartnerVerifier) EXPECT() *MockPartnerVerifierMockRecorder {
	return m.recorder
}

// GetVerifier mocks base method.
func (m *MockPartnerVerifier) GetVerifier(partnerName string) partnerverifier.Verifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVerifier", partnerName)
	ret0, _ := ret[0].(partnerverifier.Verifier)
	return ret0
}

// GetVerifier indicates an expected call of GetVerifier.
func (mr *MockPartnerVerifierMockRecorder) GetVerifier(partnerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVerifier", reflect.TypeOf((*MockPartnerVerifier)(nil).GetVerifier), partnerName)
}
