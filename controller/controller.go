package controller

// controller package provides API endpoints for the application.
// controller.go consists of common functionality used in the entire controller package.

const (
	formatValidationErrType = "format_validation_error"
	credValidationErrType   = "credential_validation_error"
	resourceNotFoundErrType = "resource_not_found_error"
	userNotFoundErrType     = "user_not_found_error"
	serviceErrType          = "service_error"
	conflictErrType         = "username_already_exists"
	DatabaseErrType         = "database_connection_error"
)
