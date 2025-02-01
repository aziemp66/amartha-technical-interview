package loan_service

import (
	loan_model "amartha-technical-interview/internal/modules/loan/model"
	"context"
)

type LoanService interface {
	CreateLoan(ctx context.Context, req loan_model.CreateLoan) (id string, err error)
	GetLoanByID(ctx context.Context, id string) (loan_model.Loan, error)
	GetAllUserLoans(ctx context.Context, userID string) ([]loan_model.Loan, error)
	GetLoanPaymentBills(ctx context.Context, id string) (float64, error)
	MakePayment(ctx context.Context, id string, amount string) error
	IsDelinquent(ctx context.Context, userID string) (bool, error)
}
