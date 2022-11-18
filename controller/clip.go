package controller

import (
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": formatValidationErrType})
		return
	}

	if err, userData.MessageID = models.InsertClip(db, userData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": serviceErrType})
		return
	}
	ctx.JSON(http.StatusCreated, userData)
}

func GETClipData(ctx *gin.Context) {
	db := models.CreateConnection()
	defer models.CloseConnection(db)
	var val int64

	if _, val = models.SelectByUsername(db, ctx.Param("username")); val == -1 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": userNotFoundErrType})
		return
	}

	clipID, _ := strconv.ParseInt(ctx.Param("clip_id"), 10, 64)
	if _, count := models.ClipCount(db, val); count <= clipID {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": resourceNotFoundErrType})
		return
	}

	_, data := models.SelectClip(db, clipID, val)
	ctx.JSON(http.StatusOK, data)
}

func GETAllClipData(ctx *gin.Context) {
	db := models.CreateConnection()
	defer models.CloseConnection(db)
	var (
		val     int64
		dataSet []models.Data
	)

	if _, val = models.SelectByUsername(db, ctx.Param("username")); val == -1 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": userNotFoundErrType})
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
	var val int64

	if _, val = models.SelectByUsername(db, ctx.Param("username")); val == -1 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": userNotFoundErrType})
		return
	}

	_ = models.DeleteClips(db, val)
	ctx.JSON(http.StatusOK, gin.H{"status": true})
}
