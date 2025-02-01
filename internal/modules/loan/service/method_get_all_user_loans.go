package loan_service

import (
	loan_model "amartha-technical-interview/internal/modules/loan/model"
	"context"

	"github.com/google/uuid"
)

func (loanService *loanService) GetAllUserLoans(ctx context.Context, userID string) ([]loan_model.Loan, error) {
	loans, err := loanService.loanRepository.GetAllUserLoans(ctx, uuid.MustParse(userID))
	if err != nil {
		return nil, err
	}

	return loans, nil
}
