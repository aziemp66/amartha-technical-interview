package loan_service

import (
	util_error "amartha-technical-interview/util/error"
	"context"
	"time"

	"github.com/google/uuid"
)

func (loanService *loanService) MakePayment(ctx context.Context, id string, amount float64) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return util_error.NewBadRequest(err, "ID is not in uuid format")
	}

	bills, err := loanService.GetLoanPaymentBills(ctx, id)
	if err != nil {
		return err
	}

	if bills != amount {
		return util_error.NewBadRequest(nil, "The amount to be paid doesnt much paid amount")
	}

	loan, err := loanService.loanRepository.GetLoanByID(ctx, parsedID)
	if err != nil {
		return err
	}

	currentTime := time.Now()
	daysPassed := int(currentTime.Sub(loan.LoanStartDate).Hours() / 24)
	weeksPassed := int(daysPassed / 7)

	if weeksPassed == loan.PaidWeeks {
		return util_error.NewBadRequest(nil, "Loans is already fully paid")
	}
	numOfInstallmentsNeedsToBePaid := weeksPassed - loan.PaidWeeks

	err = loanService.loanRepository.UpdateLoanPayment(ctx, parsedID, loan.PaidWeeks+numOfInstallmentsNeedsToBePaid)
	if err != nil {
		return err
	}

	return nil
}
