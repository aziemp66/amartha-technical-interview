package loan_http

import (
	util_http "amartha-technical-interview/util/http"

	"github.com/gin-gonic/gin"
)

func (loanHttpHandler *loanHttpHandler) IsDelinquent(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	isDelinquent, err := loanHttpHandler.loanService.IsDelinquent(ctx.Request.Context(), userID)
	if err != nil {
		ctx.Error(err)
		return
	}

	util_http.SendResponseJson(ctx, "Successfully Retrieved is delinquent status", isDelinquent)
}
