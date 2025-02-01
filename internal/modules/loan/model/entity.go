package loan_model

import (
	"time"

	"github.com/google/uuid"
)

type Loan struct {
	ID                     uuid.UUID `db:"id"`
	UserID                 uuid.UUID `db:"user_id"`
	Principal              float64   `db:"pricipal"`
	InterestRatePercentage float64   `db:"interest_rate_percentage"`
	WeeklyInstallments     int       `db:"weekly_installments"`
	PaidWeeks              int       `db:"paid_weeks"`
	LoanStartDate          time.Time `db:"loan_start_date"`
}
