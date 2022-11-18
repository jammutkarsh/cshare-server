package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/JammUtkarsh/cshare-server/models"
	"github.com/gin-gonic/gin"
)

func POSTClipData(ctx *gin.Context) {
	db := models.CreateConnection()
	defer models.CloseConnection(db)
	var (
		userData models.Data
		err      error
	)
	userData.Username = ctx.Param("username")
	if err := ctx.BindJSON(&userData); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New(formatValidationErrType))
		return
	}
	if err, userData.MessageID = models.InsertClip(db, userData); err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, errors.New(serviceErrType))
		return
	}
	ctx.JSON(http.StatusCreated, userData)
}

func GETClipData(ctx *gin.Context) {
	db := models.CreateConnection()
	defer models.CloseConnection(db)
	var val int64
	if _, val = models.SelectByUsername(db, ctx.Param("username")); val == -1 {
		_ = ctx.AbortWithError(http.StatusNotFound, errors.New(userNotFoundErrType))
		return
	}
	clipID, _ := strconv.ParseInt(ctx.Param("clip_id"), 10, 64)
	_, count := models.ClipCount(db, val)
	if count <= clipID {
		_ = ctx.AbortWithError(http.StatusNotFound, errors.New(resourceNotFoundErrType))
		return
	}
	_, data := models.SelectClip(db, clipID, val)
	ctx.JSON(http.StatusOK, data)
}

func GETAllClipData(ctx *gin.Context) {
	db := models.CreateConnection()
	defer models.CloseConnection(db)
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
	defer models.CloseConnection(db)
	_, val := models.SelectByUsername(db, ctx.Param("username"))
	if val == -1 {
		_ = ctx.AbortWithError(http.StatusUnauthorized, errors.New(userNotFoundErrType))
	}
	if err := models.DeleteClips(db, val); err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true})
}
