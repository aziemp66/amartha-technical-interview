package loan_repository_postgres

import (
	"context"
	"time"
)

func (loanRepositoryPostgres *loanRepositoryPostgres) CreateLoan(ctx context.Context, userID string, principal float64, interestRatePercentage float64, weeklyInstallments float64) (id string, err error) {
	row := loanRepositoryPostgres.db.QueryRowxContext(ctx, createLoanQuery, userID, principal, interestRatePercentage, weeklyInstallments, time.Now())

	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}
