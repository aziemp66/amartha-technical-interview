package loan_http_router

import (
	loan_http "amartha-technical-interview/internal/modules/loan/http"

	"github.com/gin-gonic/gin"
)

func BindLoanHttpRouter(router *gin.RouterGroup, handler loan_http.LoanHttpHandler) {
	router.POST("", handler.CreateLoan)
	router.POST("/payment", handler.MakePayment)
	router.GET("/:id", handler.GetLoanByID)
	router.GET("/bills/:id", handler.GetLoanPaymentBills)
	router.GET("/user/:user_id", handler.GetAllUserLoans)
	router.GET("/is-delinquent/:user_id", handler.IsDelinquent)
}
