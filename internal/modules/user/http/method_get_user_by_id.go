package user_http

import (
	user_model "amartha-technical-interview/internal/modules/user/model"
	util_error "amartha-technical-interview/util/error"
	util_http "amartha-technical-interview/util/http"

	"github.com/gin-gonic/gin"
)

func (userHttpHandler *userHttpHandler) GetUserByID(ctx *gin.Context) {
	var req user_model.GetUserIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Error(util_error.NewBadRequest(err, err.Error()))
		return
	}

	user, err := userHttpHandler.userService.GetUserByID(ctx.Request.Context(), req.ID)
	if err != nil {
		ctx.Error(err)
		return
	}

	util_http.SendResponseJson(ctx, "Successfully Retrieve User", user)
}
