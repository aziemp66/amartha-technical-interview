package loan_repository_postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
)

func (loanRepositoryPostgres *loanRepositoryPostgres) CreateLoan(ctx context.Context, userID uuid.UUID, principal float64, interestRatePercentage float64, weeklyInstallments int, loanStartDate time.Time) (id string, err error) {
	row := loanRepositoryPostgres.db.QueryRowxContext(ctx, createLoanQuery, userID, principal, interestRatePercentage, weeklyInstallments, loanStartDate)

	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}
