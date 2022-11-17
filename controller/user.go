package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/JammUtkarsh/cshare-server/auth"
	"github.com/JammUtkarsh/cshare-server/models"
	"github.com/gin-gonic/gin"
)

type ChangePassword struct {
	OldCred models.Users `json:"oldCred" binding:"required"`
	NewCred models.Users `json:"newCred" binding:"required"`
}

func POSTCreateUser(ctx *gin.Context) {
	db := models.CreateConnection()
	defer models.CloseConnection(db)
	var (
		user           models.Users
		hashedPassword string
		err            error
	)
	if err := ctx.BindJSON(&user); err != nil {
		// TODO: present error in gin.H{"error": "message"}
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New(formatValidationErrType))
		return
	}
	if err, _ := models.InsertUser(db, user.Username); err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "username already exists"})
		fmt.Println(err)
		return
	}
	if hashedPassword, err = auth.HashPassword(user); err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	if err, _ = models.InsertPasswordHash(db, user.Username, hashedPassword); err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": user.Username + " created"})
}

func POSTLogin(ctx *gin.Context) {
	db := models.CreateConnection()
	defer models.CloseConnection(db)
	var user models.Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New(formatValidationErrType))
		return
	}
	if credentialError := auth.CheckPassword(user); credentialError != nil {
		_ = ctx.AbortWithError(http.StatusUnauthorized, errors.New(credValidationErrType))
		return
	}
	tokenString, err := auth.GenerateJWT(user.Username)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func UPDATEChangePassword(ctx *gin.Context) {
	db := models.CreateConnection()
	defer models.CloseConnection(db)
	var (
		changeRequest  ChangePassword
		hashedPassword string
		err            error
	)
	if err := ctx.BindJSON(&changeRequest); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New(formatValidationErrType))
		return
	}
	if err := auth.CheckPassword(changeRequest.OldCred); err != nil {
		_ = ctx.AbortWithError(http.StatusUnauthorized, errors.New(credValidationErrType))
	}
	if hashedPassword, err = auth.HashPassword(changeRequest.NewCred); err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	if err, val := models.UpdatePassword(db, ctx.Param("username"), hashedPassword); val == false || err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, errors.New(err.Error()))
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "password changed"})
}
