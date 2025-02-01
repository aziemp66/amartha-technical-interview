package user_http_router

import (
	user_http "amartha-technical-interview/internal/modules/user/http"
	util_http_middleware "amartha-technical-interview/util/http/middleware"
	util_jwt "amartha-technical-interview/util/jwt"

	"github.com/gin-gonic/gin"
)

func BindUserHttpRouter(router *gin.RouterGroup, handler user_http.UserHttpHandler, jwtManager util_jwt.JWTManager) {
	router.GET("/:id", handler.GetUserByID)

	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)

	userRoutes := router.Group(
		"/user",
		util_http_middleware.JWTAuthentication(jwtManager),
		util_http_middleware.JWTAuthorization(util_jwt.USER_ROLE),
	)
	userRoutes.PUT("", handler.UpdateProfile)
	userRoutes.POST("/change-password", handler.ChangePassword)
}
