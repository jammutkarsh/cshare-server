package controller

import (
	"github.com/JammUtkarsh/cshare-server/auth"
	"github.com/JammUtkarsh/cshare-server/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func POSTCreateUser(ctx *gin.Context) {
	db := models.CreateConnection()
	models.CloseConnection(db)
	var user models.Users
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": formatValidationErrType})
		if err != nil {
			log.Println(err)
		}
	}
	if err, _ := models.InsertUser(db, user.Username); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
	}
	err := auth.HashPassword(user)
	if err != nil {
		log.Println(err)
	}
	token, err := auth.GenerateJWT(user.Username)
	if err == nil {
		ctx.JSON(http.StatusCreated, gin.H{"token": token})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": authErrType})
	}
}

func POSTLogin(ctx *gin.Context) {
	// get the JWT token, pasrse it and proceed.
}

func UPDATEChangePassword(ctx *gin.Context) {
	db := models.CreateConnection()
	models.CloseConnection(db)
	err := auth.CheckPassword("username", "password")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": credValidationErrType})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "password changed"})
}
