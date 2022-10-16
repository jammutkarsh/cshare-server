package controller

import (
	"errors"
	"github.com/JammUtkarsh/cshare-server/auth"
	"github.com/JammUtkarsh/cshare-server/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ChangePassword struct {
	OldCred models.Users `json:"oldCred" binding:"required"`
	NewCred models.Users `json:"newCred" binding:"required"`
}

type tokenRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func POSTCreateUser(ctx *gin.Context) {
	db := models.CreateConnection()
	models.CloseConnection(db)
	var user models.Users
	if err := ctx.BindJSON(&user); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New(formatValidationErrType))
		return
	}
	if err := auth.HashPassword(user); err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	if err, _ := models.InsertUser(db, user.Username); err != nil {
		_ = ctx.AbortWithError(http.StatusConflict, errors.New("username already exists"))
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": user.Username + " created"})
}

func POSTLogin(ctx *gin.Context) {
	db := models.CreateConnection()
	models.CloseConnection(db)
	var request tokenRequest
	var user models.Users
	if err := ctx.ShouldBindJSON(&request); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New(err.Error()))
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
	models.CloseConnection(db)
	var changeRequest ChangePassword
	// TODO: Get username and password from JSON.
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
	models.UpdatePassword(db, changeRequest.NewCred.Username, changeRequest.NewCred.Password)
	ctx.JSON(http.StatusOK, gin.H{"message": "password changed"})
}
