package loan_http

import (
	util_http "amartha-technical-interview/util/http"

	"github.com/gin-gonic/gin"
)

func (loanHttpHandler *loanHttpHandler) GetAllUserLoans(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	loans, err := loanHttpHandler.loanService.GetAllUserLoans(ctx.Request.Context(), userID)
	if err != nil {
		ctx.Error(err)
		return
	}

	util_http.SendResponseJson(ctx, "Successfully Retrieved all user's loans", loans)
}
