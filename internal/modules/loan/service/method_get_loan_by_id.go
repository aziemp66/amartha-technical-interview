package loan_service

import (
	loan_model "amartha-technical-interview/internal/modules/loan/model"
	util_error "amartha-technical-interview/util/error"
	"context"

	"github.com/google/uuid"
)

func (loanService *loanService) GetLoanByID(ctx context.Context, id string) (loan_model.Loan, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return loan_model.Loan{}, util_error.NewBadRequest(err, "ID is not in uuid format")
	}

	loan, err := loanService.loanRepository.GetLoanByID(ctx, parsedID)
	if err != nil {
		return loan_model.Loan{}, err
	}

	return loan, nil
}
