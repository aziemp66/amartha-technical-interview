package loan_model

import (
	"encoding/json"
	"time"
)

type CreateLoan struct {
	UserID                 string    `json:"user_id"`
	Principal              float64   `json:"principal"`
	InterestRatePercentage float64   `json:"interest_rate_percentage"`
	WeeklyInstallments     int       `json:"weekly_installments"`
	LoanStartDate          time.Time `json:"-"`
}

// Custom JSON unmarshaling to parse LoanStartDate from a string
func (c *CreateLoan) UnmarshalJSON(data []byte) error {
	type Alias CreateLoan
	aux := &struct {
		LoanStartDate string `json:"loan_start_date"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	parsedTime, err := time.Parse(time.RFC3339, aux.LoanStartDate)
	if err != nil {
		return err
	}

	c.LoanStartDate = parsedTime
	return nil
}

type MakePayments struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}
