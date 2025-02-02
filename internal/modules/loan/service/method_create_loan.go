package loan_service

import (
	loan_model "amartha-technical-interview/internal/modules/loan/model"
	util_error "amartha-technical-interview/util/error"
	"context"

	"github.com/google/uuid"
)

func (loanService *loanService) CreateLoan(ctx context.Context, req loan_model.CreateLoan) (id string, err error) {
	parsedID, err := uuid.Parse(req.UserID)
	if err != nil {
		return "", util_error.NewBadRequest(err, "ID is not in uuid format")
	}

	id, err = loanService.loanRepository.CreateLoan(ctx, parsedID, req.Principal, req.InterestRatePercentage, req.WeeklyInstallments, req.LoanStartDate)
	if err != nil {
		return "", err
	}

	return id, nil
}
