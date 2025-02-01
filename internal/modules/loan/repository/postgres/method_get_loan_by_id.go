package loan_repository_postgres

import (
	loan_model "backend-template/internal/modules/loan/model"
	util_error "backend-template/util/error"
	"context"
	"database/sql"

	"github.com/google/uuid"
)

func (loanRepositoryPostgres *loanRepositoryPostgres) GetLoanByID(ctx context.Context, id uuid.UUID) (loan_model.Loan, error) {
	var loan loan_model.Loan

	err := loanRepositoryPostgres.db.GetContext(ctx, &loan, getLoanByIDQuery, id)
	if err == sql.ErrNoRows {
		return loan_model.Loan{}, util_error.NewNotFound(err, "Loan is not found")
	}

	return loan, nil
}
