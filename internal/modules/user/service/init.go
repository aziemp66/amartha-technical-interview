package user_service

import (
	user_repository "amartha-technical-interview/internal/modules/user/repository"
	util_jwt "amartha-technical-interview/util/jwt"
	util_password "amartha-technical-interview/util/password"
)

type userService struct {
	userRepository  user_repository.UserRepository
	jwtManager      util_jwt.JWTManager
	passwordManager util_password.PasswordManager
}

func NewUserService(userRepository user_repository.UserRepository, jwtManager util_jwt.JWTManager, passwordManager util_password.PasswordManager) UserService {
	return &userService{
		userRepository:  userRepository,
		jwtManager:      jwtManager,
		passwordManager: passwordManager,
	}
}
