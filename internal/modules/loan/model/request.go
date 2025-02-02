package loan_model

type CreateLoan struct {
	UserID                 string  `json:"user_id"`
	Principal              float64 `json:"principal"`
	InterestRatePercentage float64 `json:"interest_rate_percentage"`
	WeeklyInstallments     float64 `json:"weekly_installments"`
}

type MakePayments struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}
