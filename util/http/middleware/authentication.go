package util_http_middleware

import (
	util_error "amartha-technical-interview/util/error"
	util_http "amartha-technical-interview/util/http"
	util_jwt "amartha-technical-interview/util/jwt"
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	BEARER = len("BEARER ")
)

func JWTAuthentication(jwtManager util_jwt.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(util_http.HEADER_AUTH)
		if authHeader == "" {
			c.Error(
				util_error.NewForbidden(errors.New("authentication header is empty"), "you are not authenticated to access this route"),
			)
			c.Abort()
			return
		} else if len(authHeader) < BEARER {
			c.Error(util_error.NewBadRequest(errors.New("authentication header invalid"), "authentication header not valid"))
			c.Abort()
			return
		}

		tokenString := authHeader[BEARER:]
		claims, err := jwtManager.VerifyAuthToken(tokenString)
		if err != nil {
			c.Error(err)
			return
		}

		c.Set("user_id", claims.ID)
		c.Set("user_role", string(claims.Role))
		c.Next()
	}
}
