package loan_repository_postgres

const (
	createLoanQuery = `INSERT INTO loan (user_id, principal, interest_rate_percentage, weekly_installments, paid_weeks, loan_start_date)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	getLoanByIDQuery = `SELECT (id, user_id, principal, interest_rate_percentage, weekly_installments, paid_weeks, loan_start_date) FROM loan WHERE
	id = $1`

	getAllUserLoan = `SELECT (id, user_id, principal, interest_rate_percentage, weekly_installments, paid_weeks, loan_start_date) FROM loan WHERE
	user_id = $1`

	UpdateLoanPayment = `UPDATE loan SET paid_weeks = $2 WHERE id = $1`
)
