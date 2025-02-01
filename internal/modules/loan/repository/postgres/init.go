package loan_repository_postgres

import (
	loan_repository "amartha-technical-interview/internal/modules/loan/repository"

	"github.com/jmoiron/sqlx"
)

type loanRepositoryPostgres struct {
	db *sqlx.DB
}

func NewLoanRepositoryPostgres(db *sqlx.DB) loan_repository.LoanRepository {
	return &loanRepositoryPostgres{
		db: db,
	}
}
