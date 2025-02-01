package user_http

import (
	user_service "amartha-technical-interview/internal/modules/user/service"
	util_jwt "amartha-technical-interview/util/jwt"
)

func NewUserHttpHandler(userService user_service.UserService, jwtManager util_jwt.JWTManager) UserHttpHandler {
	return &userHttpHandler{
		userService: userService,
		jwtManager:  jwtManager,
	}
}
