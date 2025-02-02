package loan_http

import (
	loan_model "amartha-technical-interview/internal/modules/loan/model"
	util_http "amartha-technical-interview/util/http"

	"github.com/gin-gonic/gin"
)

func (loanHttpHandler *loanHttpHandler) CreateLoan(ctx *gin.Context) {
	var req loan_model.CreateLoan
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	id, err := loanHttpHandler.loanService.CreateLoan(ctx.Request.Context(), req)
	if err != nil {
		ctx.Error(err)
		return
	}

	util_http.SendResponseJson(ctx, "Successfully Created Loan", map[string]string{
		"loan_id": id,
	})
}
