// Code generated by MockGen. DO NOT EDIT.
// Source: user_token_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	entity "github.com/education-english-web/BE-english-web/app/domain/entity"
	specifications "github.com/education-english-web/BE-english-web/app/domain/repository/specifications"
)

// MockUserTokenRepository is a mock of UserTokenRepository interface.
type MockUserTokenRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserTokenRepositoryMockRecorder
}

// MockUserTokenRepositoryMockRecorder is the mock recorder for MockUserTokenRepository.
type MockUserTokenRepositoryMockRecorder struct {
	mock *MockUserTokenRepository
}

// NewMockUserTokenRepository creates a new mock instance.
func NewMockUserTokenRepository(ctrl *gomock.Controller) *MockUserTokenRepository {
	mock := &MockUserTokenRepository{ctrl: ctrl}
	mock.recorder = &MockUserTokenRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserTokenRepository) EXPECT() *MockUserTokenRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserTokenRepository) Create(ctx context.Context, userToken *entity.UserToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserTokenRepositoryMockRecorder) Create(ctx, userToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserTokenRepository)(nil).Create), ctx, userToken)
}

// Get mocks base method.
func (m *MockUserTokenRepository) Get(ctx context.Context, specs specifications.I) (entity.UserToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, specs)
	ret0, _ := ret[0].(entity.UserToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserTokenRepositoryMockRecorder) Get(ctx, specs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserTokenRepository)(nil).Get), ctx, specs)
}

// Update mocks base method.
func (m *MockUserTokenRepository) Update(ctx context.Context, userToken *entity.UserToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, userToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserTokenRepositoryMockRecorder) Update(ctx, userToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserTokenRepository)(nil).Update), ctx, userToken)
}