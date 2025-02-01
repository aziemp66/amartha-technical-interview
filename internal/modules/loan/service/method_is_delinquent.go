package loan_service

import (
	"context"
	"time"

	"github.com/google/uuid"
)

func (loanService *loanService) IsDelinquent(ctx context.Context, userID string) (bool, error) {
	loans, err := loanService.loanRepository.GetAllUserLoans(ctx, uuid.MustParse(userID))
	if err != nil {
		return false, err
	}

	var IsDelinquent bool = false
	for _, v := range loans {
		currentTime := time.Now()
		daysPassed := int(currentTime.Sub(v.LoanStartDate).Hours() / 24)
		weeksPassed := int(daysPassed / 7)

		if weeksPassed == v.PaidWeeks {
			continue
		}
		numOfInstallmentsNeedsToBePaid := weeksPassed - v.PaidWeeks

		if numOfInstallmentsNeedsToBePaid >= 2 {
			IsDelinquent = true
			break
		}
	}

	return IsDelinquent, nil
}
