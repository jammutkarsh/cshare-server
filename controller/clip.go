package controller

import (
	"errors"
	"github.com/JammUtkarsh/cshare-server/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func POSTClipData(ctx *gin.Context) {
	db := models.CreateConnection()
	models.CloseConnection(db)
	var userData models.Data
	if err := ctx.BindJSON(&userData); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New(formatValidationErrType))
		log.Println(err)
		return
	}
	clip := models.Data{
		Username: userData.Username,
		Message:  userData.Message,
		Secret:   userData.Secret,
	}
	if err, _ := models.InsertClip(db, clip); err != nil {
		_ = ctx.AbortWithError(http.StatusNotFound, errors.New(serviceErrType))
		return
	}
	ctx.JSON(http.StatusCreated, &userData)

}

func GETClipData(ctx *gin.Context) {
	db := models.CreateConnection()
	models.CloseConnection(db)
	err, val := models.SelectByUsername(db, ctx.Param("username"))
	if err != nil {
		log.Println(err)
	}
	if val == -1 {
		_ = ctx.AbortWithError(http.StatusNotFound, errors.New(userNotFoundErrType))
		return
	}
	clipID, _ := strconv.ParseInt(ctx.Param("clip_id"), 10, 64)
	_, count := models.ClipCount(db, val)
	if count <= clipID {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New(resourceNotFoundErrType))
	}
	_, data := models.SelectClip(db, clipID, val)
	ctx.SecureJSON(http.StatusOK, data)
}

func GETAllClipData(ctx *gin.Context) {
	db := models.CreateConnection()
	models.CloseConnection(db)
	var dataSet []models.Data
	_, val := models.SelectByUsername(db, ctx.Param("username"))
	if val == -1 {
		_ = ctx.AbortWithError(http.StatusUnauthorized, errors.New(userNotFoundErrType))
		return
	}
	_, count := models.ClipCount(db, val)
	for i := int64(1); i <= count; i++ {
		_, data := models.SelectClip(db, i, val)
		dataSet = append(dataSet, data)
	}
	ctx.JSON(http.StatusOK, dataSet)
}

func DELETEAllClipData(ctx *gin.Context) {
	db := models.CreateConnection()
	models.CloseConnection(db)
	_, val := models.SelectByUsername(db, ctx.Param("username"))
	if val == -1 {
		_ = ctx.AbortWithError(http.StatusUnauthorized, errors.New(userNotFoundErrType))
	}
	_, count := models.ClipCount(db, val)
	for i := int64(0); i <= count; i++ {
		err := models.DeleteClip(db, i, val)
		if err != nil {
			log.Println(err)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true})
}
