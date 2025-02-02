package loan_http

import (
	"github.com/gin-gonic/gin"
)

type LoanHttpHandler interface {
	CreateLoan(ctx *gin.Context)
	GetLoanByID(ctx *gin.Context)
	GetAllUserLoans(ctx *gin.Context)
	GetLoanPaymentBills(ctx *gin.Context)
	MakePayment(ctx *gin.Context)
	IsDelinquent(ctx *gin.Context)
}
