// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nkolosov/whip-round/internal/repository (interfaces: User,Session)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	domain "github.com/nkolosov/whip-round/internal/domain"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUser) CreateUser(arg0 context.Context, arg1 *domain.User) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUser)(nil).CreateUser), arg0, arg1)
}

// GetUserByEmail mocks base method.
func (m *MockUser) GetUserByEmail(arg0 context.Context, arg1 string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserMockRecorder) GetUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUser)(nil).GetUserByEmail), arg0, arg1)
}

// MockSession is a mock of Session interface.
type MockSession struct {
	ctrl     *gomock.Controller
	recorder *MockSessionMockRecorder
}

// MockSessionMockRecorder is the mock recorder for MockSession.
type MockSessionMockRecorder struct {
	mock *MockSession
}

// NewMockSession creates a new mock instance.
func NewMockSession(ctrl *gomock.Controller) *MockSession {
	mock := &MockSession{ctrl: ctrl}
	mock.recorder = &MockSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSession) EXPECT() *MockSessionMockRecorder {
	return m.recorder
}

// CreateRefreshToken mocks base method.
func (m *MockSession) CreateRefreshToken(arg0 context.Context, arg1 *domain.RefreshSession) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRefreshToken", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRefreshToken indicates an expected call of CreateRefreshToken.
func (mr *MockSessionMockRecorder) CreateRefreshToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRefreshToken", reflect.TypeOf((*MockSession)(nil).CreateRefreshToken), arg0, arg1)
}

// DeleteRefreshToken mocks base method.
func (m *MockSession) DeleteRefreshToken(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRefreshToken", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRefreshToken indicates an expected call of DeleteRefreshToken.
func (mr *MockSessionMockRecorder) DeleteRefreshToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRefreshToken", reflect.TypeOf((*MockSession)(nil).DeleteRefreshToken), arg0, arg1)
}

// DeleteRefreshTokenByUserId mocks base method.
func (m *MockSession) DeleteRefreshTokenByUserId(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRefreshTokenByUserId", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRefreshTokenByUserId indicates an expected call of DeleteRefreshTokenByUserId.
func (mr *MockSessionMockRecorder) DeleteRefreshTokenByUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRefreshTokenByUserId", reflect.TypeOf((*MockSession)(nil).DeleteRefreshTokenByUserId), arg0, arg1)
}

// GetRefreshToken mocks base method.
func (m *MockSession) GetRefreshToken(arg0 context.Context, arg1 string) (*domain.RefreshSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRefreshToken", arg0, arg1)
	ret0, _ := ret[0].(*domain.RefreshSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRefreshToken indicates an expected call of GetRefreshToken.
func (mr *MockSessionMockRecorder) GetRefreshToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRefreshToken", reflect.TypeOf((*MockSession)(nil).GetRefreshToken), arg0, arg1)
}
