package controller

// user.go consists of methods concerning user endpoints;
// It provides a starting point for the clip endpoints.
// Every method follows a standard procedure of
// 1. JSON validation.
// 2. Credential validation.
// 3. Database operations.
// 4. Returning a response.

import (
	"net/http"

	"github.com/JammUtkarsh/cshare-server/auth"
	"github.com/JammUtkarsh/cshare-server/models"
	"github.com/gin-gonic/gin"
)

type changePassword struct {
	OldCred models.Users `json:"oldCred" binding:"required"`
	NewCred models.Users `json:"newCred" binding:"required"`
}

// POSTCreateUser is POST HTTP method; accepts a user entry in the database for a given valid JSON.
func POSTCreateUser(ctx *gin.Context) {
	db := models.CreateConnection()
	defer models.CloseConnection(db)
	var (
		user           models.Users
		hashedPassword string
		err            error
	)

	if err = ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": formatValidationErrType})
		return
	}

	if err, _ = models.InsertUser(db, user.Username); err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": conflictErrType})
		return
	}

	if hashedPassword, err = auth.HashPassword(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": serviceErrType})
		return
	}

	if err = models.InsertPasswordHash(db, user.Username, hashedPassword); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": serviceErrType})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": user.Username + " created"})
}

// POSTLogin is POST HTTP method, validates credentials of an existing user and returns a JWT.
func POSTLogin(ctx *gin.Context) {
	db := models.CreateConnection()
	defer models.CloseConnection(db)
	var (
		user        models.Users
		tokenString string
		err         error
	)

	if err = ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": formatValidationErrType})
		return
	}

	if err = auth.CheckPassword(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": credValidationErrType})
		return
	}

	if tokenString, err = auth.GenerateJWT(user.Username); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": serviceErrType})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}
