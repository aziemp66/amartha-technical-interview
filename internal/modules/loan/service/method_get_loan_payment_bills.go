package loan_service

import (
	util_error "amartha-technical-interview/util/error"
	"context"
	"time"

	"github.com/google/uuid"
)

func (loanService *loanService) GetLoanPaymentBills(ctx context.Context, id string) (float64, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return 0, util_error.NewBadRequest(err, "ID is not in uuid format")
	}

	loan, err := loanService.loanRepository.GetLoanByID(ctx, parsedID)
	if err != nil {
		return 0, err
	}

	currentTime := time.Now()
	daysPassed := int(currentTime.Sub(loan.LoanStartDate).Hours() / 24)
	weeksPassed := int(daysPassed / 7)

	if weeksPassed == loan.PaidWeeks {
		return 0, nil
	}
	numOfInstallmentsNeedsToBePaid := weeksPassed - loan.PaidWeeks

	paymentAmounts := float64(numOfInstallmentsNeedsToBePaid) * ((loan.Principal + (loan.Principal * loan.InterestRatePercentage)) / float64(loan.WeeklyInstallments))

	return paymentAmounts, nil
}
