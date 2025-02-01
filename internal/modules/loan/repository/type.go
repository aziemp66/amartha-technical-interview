package loan_repository

import (
	loan_model "amartha-technical-interview/internal/modules/loan/model"
	"context"

	"github.com/google/uuid"
)

type LoanRepository interface {
	CreateLoan(ctx context.Context, userID string, principal, interestRatePercentage, weeklyInstallments float64) (id string, err error)
	GetLoanByID(ctx context.Context, id uuid.UUID) (loan_model.Loan, error)
	GetAllUserLoans(ctx context.Context, userID uuid.UUID) ([]loan_model.Loan, error)
	UpdateLoanPayment(ctx context.Context, id uuid.UUID, weeksPayed int) error
}
