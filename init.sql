CREATE TABLE users (
	id UUID PRIMARY KEY,
	email VARCHAR(255) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL,
	name VARCHAR(255) NOT NULL,
	address TEXT
);

CREATE TABLE loans (
	id UUID PRIMARY KEY,
	user_id UUID NOT NULL,
	principal DOUBLE PRECISION NOT NULL,
	interest_rate_percentage DOUBLE PRECISION NOT NULL,
	weekly_installments INT NOT NULL,
	paid_weeks INT NOT NULL DEFAULT 0,
	loan_start_date TIMESTAMP NOT NULL,
	CONSTRAINT fk_loans_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);