// Code generated by MockGen. DO NOT EDIT.
// Source: internal/modules/user/repository/type.go
//
// Generated by this command:
//
//	mockgen -package=mock_repository -source=internal/modules/user/repository/type.go -destination=mock/repository/user_repository_mock.go -typed=true
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	user_model "backend-template/internal/modules/user/model"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// ChangePassword mocks base method.
func (m *MockUserRepository) ChangePassword(ctx context.Context, id, hashedPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", ctx, id, hashedPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserRepositoryMockRecorder) ChangePassword(ctx, id, hashedPassword any) *MockUserRepositoryChangePasswordCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserRepository)(nil).ChangePassword), ctx, id, hashedPassword)
	return &MockUserRepositoryChangePasswordCall{Call: call}
}

// MockUserRepositoryChangePasswordCall wrap *gomock.Call
type MockUserRepositoryChangePasswordCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserRepositoryChangePasswordCall) Return(err error) *MockUserRepositoryChangePasswordCall {
	c.Call = c.Call.Return(err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserRepositoryChangePasswordCall) Do(f func(context.Context, string, string) error) *MockUserRepositoryChangePasswordCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserRepositoryChangePasswordCall) DoAndReturn(f func(context.Context, string, string) error) *MockUserRepositoryChangePasswordCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(ctx context.Context, email, hashedPassword, name, address string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, email, hashedPassword, name, address)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(ctx, email, hashedPassword, name, address any) *MockUserRepositoryCreateUserCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), ctx, email, hashedPassword, name, address)
	return &MockUserRepositoryCreateUserCall{Call: call}
}

// MockUserRepositoryCreateUserCall wrap *gomock.Call
type MockUserRepositoryCreateUserCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserRepositoryCreateUserCall) Return(id string, err error) *MockUserRepositoryCreateUserCall {
	c.Call = c.Call.Return(id, err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserRepositoryCreateUserCall) Do(f func(context.Context, string, string, string, string) (string, error)) *MockUserRepositoryCreateUserCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserRepositoryCreateUserCall) DoAndReturn(f func(context.Context, string, string, string, string) (string, error)) *MockUserRepositoryCreateUserCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteUser mocks base method.
func (m *MockUserRepository) DeleteUser(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepositoryMockRecorder) DeleteUser(ctx, id any) *MockUserRepositoryDeleteUserCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepository)(nil).DeleteUser), ctx, id)
	return &MockUserRepositoryDeleteUserCall{Call: call}
}

// MockUserRepositoryDeleteUserCall wrap *gomock.Call
type MockUserRepositoryDeleteUserCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserRepositoryDeleteUserCall) Return(err error) *MockUserRepositoryDeleteUserCall {
	c.Call = c.Call.Return(err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserRepositoryDeleteUserCall) Do(f func(context.Context, string) error) *MockUserRepositoryDeleteUserCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserRepositoryDeleteUserCall) DoAndReturn(f func(context.Context, string) error) *MockUserRepositoryDeleteUserCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetUserByEmail mocks base method.
func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (user_model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(user_model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserRepositoryMockRecorder) GetUserByEmail(ctx, email any) *MockUserRepositoryGetUserByEmailCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).GetUserByEmail), ctx, email)
	return &MockUserRepositoryGetUserByEmailCall{Call: call}
}

// MockUserRepositoryGetUserByEmailCall wrap *gomock.Call
type MockUserRepositoryGetUserByEmailCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserRepositoryGetUserByEmailCall) Return(res user_model.User, err error) *MockUserRepositoryGetUserByEmailCall {
	c.Call = c.Call.Return(res, err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserRepositoryGetUserByEmailCall) Do(f func(context.Context, string) (user_model.User, error)) *MockUserRepositoryGetUserByEmailCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserRepositoryGetUserByEmailCall) DoAndReturn(f func(context.Context, string) (user_model.User, error)) *MockUserRepositoryGetUserByEmailCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetUserByID mocks base method.
func (m *MockUserRepository) GetUserByID(ctx context.Context, id string) (user_model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(user_model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserRepositoryMockRecorder) GetUserByID(ctx, id any) *MockUserRepositoryGetUserByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserRepository)(nil).GetUserByID), ctx, id)
	return &MockUserRepositoryGetUserByIDCall{Call: call}
}

// MockUserRepositoryGetUserByIDCall wrap *gomock.Call
type MockUserRepositoryGetUserByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserRepositoryGetUserByIDCall) Return(res user_model.User, err error) *MockUserRepositoryGetUserByIDCall {
	c.Call = c.Call.Return(res, err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserRepositoryGetUserByIDCall) Do(f func(context.Context, string) (user_model.User, error)) *MockUserRepositoryGetUserByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserRepositoryGetUserByIDCall) DoAndReturn(f func(context.Context, string) (user_model.User, error)) *MockUserRepositoryGetUserByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateUser mocks base method.
func (m *MockUserRepository) UpdateUser(ctx context.Context, id, name, address string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, id, name, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserRepositoryMockRecorder) UpdateUser(ctx, id, name, address any) *MockUserRepositoryUpdateUserCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepository)(nil).UpdateUser), ctx, id, name, address)
	return &MockUserRepositoryUpdateUserCall{Call: call}
}

// MockUserRepositoryUpdateUserCall wrap *gomock.Call
type MockUserRepositoryUpdateUserCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserRepositoryUpdateUserCall) Return(err error) *MockUserRepositoryUpdateUserCall {
	c.Call = c.Call.Return(err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserRepositoryUpdateUserCall) Do(f func(context.Context, string, string, string) error) *MockUserRepositoryUpdateUserCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserRepositoryUpdateUserCall) DoAndReturn(f func(context.Context, string, string, string) error) *MockUserRepositoryUpdateUserCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// VerifyUser mocks base method.
func (m *MockUserRepository) VerifyUser(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyUser indicates an expected call of VerifyUser.
func (mr *MockUserRepositoryMockRecorder) VerifyUser(ctx, id any) *MockUserRepositoryVerifyUserCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUser", reflect.TypeOf((*MockUserRepository)(nil).VerifyUser), ctx, id)
	return &MockUserRepositoryVerifyUserCall{Call: call}
}

// MockUserRepositoryVerifyUserCall wrap *gomock.Call
type MockUserRepositoryVerifyUserCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserRepositoryVerifyUserCall) Return(err error) *MockUserRepositoryVerifyUserCall {
	c.Call = c.Call.Return(err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserRepositoryVerifyUserCall) Do(f func(context.Context, string) error) *MockUserRepositoryVerifyUserCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserRepositoryVerifyUserCall) DoAndReturn(f func(context.Context, string) error) *MockUserRepositoryVerifyUserCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
