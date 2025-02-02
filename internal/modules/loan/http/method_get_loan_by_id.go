package loan_http

import (
	util_http "amartha-technical-interview/util/http"

	"github.com/gin-gonic/gin"
)

func (loanHttpHandler *loanHttpHandler) GetLoanByID(ctx *gin.Context) {
	id := ctx.Param("id")

	loan, err := loanHttpHandler.loanService.GetLoanByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.Error(err)
		return
	}

	util_http.SendResponseJson(ctx, "Successfully Retrieved Loan", loan)
}
