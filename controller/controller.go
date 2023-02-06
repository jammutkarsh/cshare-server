package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// controller package provides API endpoints for the application.
// controller.go consists of common functionality used in the entire controller package.

const (
	formatValidationErrType = "format_validation_error"
	credValidationErrType   = "credential_validation_error"
	resourceNotFoundErrType = "resource_not_found_error"
	userNotFoundErrType     = "user_not_found_error"
	serviceErrType          = "service_error"
	conflictErrType         = "username_already_exists"
	databaseErrType         = "database_connection_error"
)

// HomepageHandler is a GET HTTP method; returns a welcome message.
func HomepageHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message":"Welcome to cShare, a clipboard sharing service. Visit https://github.com/JammUtkarsh/cshare-server for more info"})
	
}