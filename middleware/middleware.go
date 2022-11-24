package middleware

// middleware package currently deals with authorization of user to access resource endpoints.

import (
	"errors"
	"strings"

	"github.com/JammUtkarsh/cshare-server/auth"
	"github.com/gin-gonic/gin"
)

const authErrType = "authentication_error"

// Auth verifies a user's authenticity for a given JWT string from Authorization header.
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		var tokenString string
		if tokenString = context.GetHeader("Authorization"); tokenString == "" {
			_ = context.AbortWithError(401, errors.New("request_does_not_contain_an_access_token"))
			return
		}
		// removes the word 'Bearer' from the `Authorization` header to process a valid JWT string.
		tokenString = strings.Split(tokenString, "Bearer ")[1]
		if err := auth.ValidateToken(tokenString); err != nil {
			_ = context.AbortWithError(401, errors.New(authErrType))
			return
		}
		context.Next()
	}
}
