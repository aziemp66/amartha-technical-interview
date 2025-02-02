package init_app

import (
	loan_http "amartha-technical-interview/internal/modules/loan/http"
	loan_http_router "amartha-technical-interview/internal/modules/loan/http/router"
	loan_repository_postgres "amartha-technical-interview/internal/modules/loan/repository/postgres"
	loan_service "amartha-technical-interview/internal/modules/loan/service"
	user_http "amartha-technical-interview/internal/modules/user/http"
	user_http_router "amartha-technical-interview/internal/modules/user/http/router"
	user_repository_postgres "amartha-technical-interview/internal/modules/user/repository/postgres"
	user_service "amartha-technical-interview/internal/modules/user/service"
	util_jwt "amartha-technical-interview/util/jwt"
	util_password "amartha-technical-interview/util/password"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitializeApp(router *gin.Engine, db *sqlx.DB, jwtManager util_jwt.JWTManager, passwordManager util_password.PasswordManager) {
	userRepository := user_repository_postgres.NewUserRepositoryPostgres(db)
	userService := user_service.NewUserService(userRepository, jwtManager, passwordManager)
	userHandler := user_http.NewUserHttpHandler(userService, jwtManager)

	user_http_router.BindUserHttpRouter(router.Group("/user"), userHandler, jwtManager)

	loanRepository := loan_repository_postgres.NewLoanRepositoryPostgres(db)
	loanService := loan_service.NewLoanService(loanRepository)
	loanHandler := loan_http.NewLoanHttpHandler(loanService, jwtManager)

	loan_http_router.BindLoanHttpRouter(router.Group("/loan"), loanHandler)
}
