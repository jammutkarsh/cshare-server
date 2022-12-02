package middleware

// middleware package currently deals with authorization of user to access resource endpoints.

import (
	"strings"

	"github.com/JammUtkarsh/cshare-server/auth"
	"github.com/gin-gonic/gin"
)

const (
	authErrType  = "authentication_error"
	missingToken = "request_does_not_contain_an_access_token"
)

// Auth verifies a user's authenticity for a given JWT string from Authorization header.
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tokenString string
		if tokenString = ctx.GetHeader("Authorization"); tokenString == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": missingToken})
			return
		}
		// removes the word 'Bearer' from the `Authorization` header to process a valid JWT string.
		tokenString = strings.Split(tokenString, "Bearer ")[1]
		if err := auth.ValidateToken(tokenString); err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": authErrType})
			return
		}
		ctx.Next()
	}
}
