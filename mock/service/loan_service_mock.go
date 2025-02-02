// Code generated by MockGen. DO NOT EDIT.
// Source: internal/modules/loan/service/type.go
//
// Generated by this command:
//
//	mockgen -package=mock_service -source=internal/modules/loan/service/type.go -destination=mock/service/loan_service_mock.go -typed
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	loan_model "amartha-technical-interview/internal/modules/loan/model"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockLoanService is a mock of LoanService interface.
type MockLoanService struct {
	ctrl     *gomock.Controller
	recorder *MockLoanServiceMockRecorder
	isgomock struct{}
}

// MockLoanServiceMockRecorder is the mock recorder for MockLoanService.
type MockLoanServiceMockRecorder struct {
	mock *MockLoanService
}

// NewMockLoanService creates a new mock instance.
func NewMockLoanService(ctrl *gomock.Controller) *MockLoanService {
	mock := &MockLoanService{ctrl: ctrl}
	mock.recorder = &MockLoanServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoanService) EXPECT() *MockLoanServiceMockRecorder {
	return m.recorder
}

// CreateLoan mocks base method.
func (m *MockLoanService) CreateLoan(ctx context.Context, req loan_model.CreateLoan) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLoan", ctx, req)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLoan indicates an expected call of CreateLoan.
func (mr *MockLoanServiceMockRecorder) CreateLoan(ctx, req any) *MockLoanServiceCreateLoanCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLoan", reflect.TypeOf((*MockLoanService)(nil).CreateLoan), ctx, req)
	return &MockLoanServiceCreateLoanCall{Call: call}
}

// MockLoanServiceCreateLoanCall wrap *gomock.Call
type MockLoanServiceCreateLoanCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoanServiceCreateLoanCall) Return(id string, err error) *MockLoanServiceCreateLoanCall {
	c.Call = c.Call.Return(id, err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoanServiceCreateLoanCall) Do(f func(context.Context, loan_model.CreateLoan) (string, error)) *MockLoanServiceCreateLoanCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoanServiceCreateLoanCall) DoAndReturn(f func(context.Context, loan_model.CreateLoan) (string, error)) *MockLoanServiceCreateLoanCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetAllUserLoans mocks base method.
func (m *MockLoanService) GetAllUserLoans(ctx context.Context, userID string) ([]loan_model.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserLoans", ctx, userID)
	ret0, _ := ret[0].([]loan_model.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserLoans indicates an expected call of GetAllUserLoans.
func (mr *MockLoanServiceMockRecorder) GetAllUserLoans(ctx, userID any) *MockLoanServiceGetAllUserLoansCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserLoans", reflect.TypeOf((*MockLoanService)(nil).GetAllUserLoans), ctx, userID)
	return &MockLoanServiceGetAllUserLoansCall{Call: call}
}

// MockLoanServiceGetAllUserLoansCall wrap *gomock.Call
type MockLoanServiceGetAllUserLoansCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoanServiceGetAllUserLoansCall) Return(arg0 []loan_model.Loan, arg1 error) *MockLoanServiceGetAllUserLoansCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoanServiceGetAllUserLoansCall) Do(f func(context.Context, string) ([]loan_model.Loan, error)) *MockLoanServiceGetAllUserLoansCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoanServiceGetAllUserLoansCall) DoAndReturn(f func(context.Context, string) ([]loan_model.Loan, error)) *MockLoanServiceGetAllUserLoansCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetLoanByID mocks base method.
func (m *MockLoanService) GetLoanByID(ctx context.Context, id string) (loan_model.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoanByID", ctx, id)
	ret0, _ := ret[0].(loan_model.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoanByID indicates an expected call of GetLoanByID.
func (mr *MockLoanServiceMockRecorder) GetLoanByID(ctx, id any) *MockLoanServiceGetLoanByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoanByID", reflect.TypeOf((*MockLoanService)(nil).GetLoanByID), ctx, id)
	return &MockLoanServiceGetLoanByIDCall{Call: call}
}

// MockLoanServiceGetLoanByIDCall wrap *gomock.Call
type MockLoanServiceGetLoanByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoanServiceGetLoanByIDCall) Return(arg0 loan_model.Loan, arg1 error) *MockLoanServiceGetLoanByIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoanServiceGetLoanByIDCall) Do(f func(context.Context, string) (loan_model.Loan, error)) *MockLoanServiceGetLoanByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoanServiceGetLoanByIDCall) DoAndReturn(f func(context.Context, string) (loan_model.Loan, error)) *MockLoanServiceGetLoanByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetLoanPaymentBills mocks base method.
func (m *MockLoanService) GetLoanPaymentBills(ctx context.Context, id string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoanPaymentBills", ctx, id)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoanPaymentBills indicates an expected call of GetLoanPaymentBills.
func (mr *MockLoanServiceMockRecorder) GetLoanPaymentBills(ctx, id any) *MockLoanServiceGetLoanPaymentBillsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoanPaymentBills", reflect.TypeOf((*MockLoanService)(nil).GetLoanPaymentBills), ctx, id)
	return &MockLoanServiceGetLoanPaymentBillsCall{Call: call}
}

// MockLoanServiceGetLoanPaymentBillsCall wrap *gomock.Call
type MockLoanServiceGetLoanPaymentBillsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoanServiceGetLoanPaymentBillsCall) Return(arg0 float64, arg1 error) *MockLoanServiceGetLoanPaymentBillsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoanServiceGetLoanPaymentBillsCall) Do(f func(context.Context, string) (float64, error)) *MockLoanServiceGetLoanPaymentBillsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoanServiceGetLoanPaymentBillsCall) DoAndReturn(f func(context.Context, string) (float64, error)) *MockLoanServiceGetLoanPaymentBillsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsDelinquent mocks base method.
func (m *MockLoanService) IsDelinquent(ctx context.Context, userID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDelinquent", ctx, userID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsDelinquent indicates an expected call of IsDelinquent.
func (mr *MockLoanServiceMockRecorder) IsDelinquent(ctx, userID any) *MockLoanServiceIsDelinquentCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDelinquent", reflect.TypeOf((*MockLoanService)(nil).IsDelinquent), ctx, userID)
	return &MockLoanServiceIsDelinquentCall{Call: call}
}

// MockLoanServiceIsDelinquentCall wrap *gomock.Call
type MockLoanServiceIsDelinquentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoanServiceIsDelinquentCall) Return(arg0 bool, arg1 error) *MockLoanServiceIsDelinquentCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoanServiceIsDelinquentCall) Do(f func(context.Context, string) (bool, error)) *MockLoanServiceIsDelinquentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoanServiceIsDelinquentCall) DoAndReturn(f func(context.Context, string) (bool, error)) *MockLoanServiceIsDelinquentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MakePayment mocks base method.
func (m *MockLoanService) MakePayment(ctx context.Context, id string, amount float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakePayment", ctx, id, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakePayment indicates an expected call of MakePayment.
func (mr *MockLoanServiceMockRecorder) MakePayment(ctx, id, amount any) *MockLoanServiceMakePaymentCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakePayment", reflect.TypeOf((*MockLoanService)(nil).MakePayment), ctx, id, amount)
	return &MockLoanServiceMakePaymentCall{Call: call}
}

// MockLoanServiceMakePaymentCall wrap *gomock.Call
type MockLoanServiceMakePaymentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoanServiceMakePaymentCall) Return(arg0 error) *MockLoanServiceMakePaymentCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoanServiceMakePaymentCall) Do(f func(context.Context, string, float64) error) *MockLoanServiceMakePaymentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoanServiceMakePaymentCall) DoAndReturn(f func(context.Context, string, float64) error) *MockLoanServiceMakePaymentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
