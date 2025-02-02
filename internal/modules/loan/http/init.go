package loan_http

import (
	loan_service "amartha-technical-interview/internal/modules/loan/service"
	util_jwt "amartha-technical-interview/util/jwt"
)

type loanHttpHandler struct {
	loanService loan_service.LoanService
	jwtManager  util_jwt.JWTManager
}

func NewLoanHttpHandler(loanService loan_service.LoanService, jwtManager util_jwt.JWTManager) LoanHttpHandler {
	return &loanHttpHandler{
		loanService: loanService,
		jwtManager:  jwtManager,
	}
}
