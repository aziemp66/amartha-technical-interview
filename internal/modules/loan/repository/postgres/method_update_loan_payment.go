package loan_repository_postgres

import (
	"context"

	"github.com/google/uuid"
)

func (loanRepositoryPostgres *loanRepositoryPostgres) UpdateLoanPayment(ctx context.Context, id uuid.UUID, weeksPayed int) error {
	_, err := loanRepositoryPostgres.db.ExecContext(ctx, UpdateLoanPayment, id, weeksPayed)
	if err != nil {
		return err
	}

	return nil
}
