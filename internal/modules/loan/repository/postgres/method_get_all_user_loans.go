package loan_repository_postgres

import (
	loan_model "backend-template/internal/modules/loan/model"
	"context"

	"github.com/google/uuid"
)

func (loanRepositoryPostgres *loanRepositoryPostgres) GetAllUserLoans(ctx context.Context, userID uuid.UUID) ([]loan_model.Loan, error) {
	var loans []loan_model.Loan

	err := loanRepositoryPostgres.db.SelectContext(ctx, &loans, getAllUserLoan)
	if err != nil {
		return nil, err
	}

	return loans, nil
}
