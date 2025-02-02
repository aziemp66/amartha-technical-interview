package loan_http

import (
	loan_model "amartha-technical-interview/internal/modules/loan/model"
	util_http "amartha-technical-interview/util/http"

	"github.com/gin-gonic/gin"
)

func (loanHttpHandler *loanHttpHandler) MakePayment(ctx *gin.Context) {
	var req loan_model.MakePayments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	err := loanHttpHandler.loanService.MakePayment(ctx.Request.Context(), req.ID, req.Amount)
	if err != nil {
		ctx.Error(err)
		return
	}

	util_http.SendResponseJson(ctx, "Successfully made payment", nil)
}
