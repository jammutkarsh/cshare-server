package controller

import (
	"errors"
	"log"
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
	var user models.Users
	if err := ctx.BindJSON(&user); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New(formatValidationErrType))
		return
	}
	if err, _ := models.InsertUser(db, user.Username); err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "username already exists"})
		return
	}
	if err := auth.HashPassword(user); err != nil {
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
	var changeRequest ChangePassword
	if err := ctx.BindJSON(&changeRequest); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New(formatValidationErrType))
		log.Println(err)
		return
	}
	if err := auth.CheckPassword(changeRequest.OldCred); err != nil {
		_ = ctx.AbortWithError(http.StatusUnauthorized, errors.New(credValidationErrType))
	}
	if err := auth.HashPassword(changeRequest.NewCred); err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	if err, val := models.UpdatePassword(db, changeRequest.NewCred.Username, changeRequest.NewCred.Password); val == false || err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, errors.New(err.Error()))
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "password changed"})
}
