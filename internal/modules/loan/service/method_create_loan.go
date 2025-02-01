package loan_service

import (
	loan_model "amartha-technical-interview/internal/modules/loan/model"
	"context"

	"github.com/google/uuid"
)

func (loanService *loanService) CreateLoan(ctx context.Context, userID string, req loan_model.CreateLoan) (id string, err error) {
	id, err = loanService.loanRepository.CreateLoan(ctx, uuid.MustParse(userID), req.Principal, req.InterestRatePercentage, req.WeeklyInstallments)
	if err != nil {
		return "", err
	}

	return id, nil
}
