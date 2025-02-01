package loan_service

import loan_repository "amartha-technical-interview/internal/modules/loan/repository"

type loanService struct {
	loanRepository loan_repository.LoanRepository
}

func NewLoanService(loanRepository loan_repository.LoanRepository) LoanService {
	return &loanService{
		loanRepository: loanRepository,
	}
}
