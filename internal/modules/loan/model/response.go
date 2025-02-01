package loan_model

import (
	"encoding/json"
	"time"
)

type LoanResponse struct {
	ID                     string    `json:"id"`
	UserID                 string    `json:"user_id"`
	Principal              float64   `json:"pricipal"`
	InterestRatePercentage float64   `json:"interest_rate_percentage"`
	WeeklyInstallments     int       `json:"weekly_installments"`
	PaidWeeks              int       `json:"paid_weeks"`
	LoanStartDate          time.Time `json:"-"`
}

// Custom JSON Marshaling to format LoanStartDate in RFC3339
func (l LoanResponse) MarshalJSON() ([]byte, error) {
	type Alias LoanResponse
	return json.Marshal(&struct {
		LoanStartDate string `json:"loan_start_date"`
		*Alias
	}{
		LoanStartDate: l.LoanStartDate.Format(time.RFC3339),
		Alias:         (*Alias)(&l),
	})
}
