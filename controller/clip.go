package controller

// clips.go consists of methods concerning clip API endpoints;
// They are secured by a JWT which are generated at user API endpoint.
// Every method follows a standard procedure of
// 1. JSON validation.
// 2. Database operations.
// 3. Returning a response.

import (
	"net/http"
	"strconv"

	"github.com/JammUtkarsh/cshare-server/models"
	"github.com/gin-gonic/gin"
)

// POSTClipData is POST HTTP method; For a given user with a valid clips JSON and stores it in DB.
// returns appropriate response with status code.
func POSTClipData(ctx *gin.Context) {
	var userData models.Data
	userData.Username = ctx.Param("username")

	if err := ctx.BindJSON(&userData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": formatValidationErrType})
		return
	}
	
	db, err := models.CreateConnection()
	defer models.CloseConnection(db)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": databaseErrType})
		return
	}

	if userData.MessageID, err = models.InsertClip(db, userData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": serviceErrType})
		return
	}
	ctx.JSON(http.StatusCreated, userData)
}

// GETClipData is a GET HTTP method; returns a `single clip` JSON data for a given user and messageID.
func GETClipData(ctx *gin.Context) {
	db, err := models.CreateConnection()
	defer models.CloseConnection(db)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": databaseErrType})
		return
	}

	var val int64

	if val, _ = models.GetUserID(db, ctx.Param("username")); val == -1 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": userNotFoundErrType})
		return
	}

	clipID, _ := strconv.ParseInt(ctx.Param("clip_id"), 10, 64)
	if count, _ := models.ClipCount(db, val); clipID > count && clipID > 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": resourceNotFoundErrType})
		return
	}

	data, _ := models.SelectClip(db, clipID, val)
	ctx.JSON(http.StatusOK, data)
}

// GETAllClipData is a GET HTTP method; returns `all clips` JSON data for a given user.
func GETAllClipData(ctx *gin.Context) {
	db, err := models.CreateConnection()
	defer models.CloseConnection(db)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": databaseErrType})
		return
	}

	var (
		val     int64
		dataSet []models.Data
	)

	if val, _ = models.GetUserID(db, ctx.Param("username")); val == -1 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": userNotFoundErrType})
		return
	}

	count, _ := models.ClipCount(db, val)
	for i := int64(1); i <= count; i++ {
		data, _ := models.SelectClip(db, i, val)
		dataSet = append(dataSet, data)
	}
	ctx.JSON(http.StatusOK, dataSet)
}

// DELETEAllClipData is a DELETE HTTP method; deletes all the clips for a given user.
func DELETEAllClipData(ctx *gin.Context) {
	db, err := models.CreateConnection()
	defer models.CloseConnection(db)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": databaseErrType})
		return
	}

	var val int64

	if val, _ = models.GetUserID(db, ctx.Param("username")); val == -1 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": userNotFoundErrType})
		return
	}

	_ = models.DeleteClips(db, val)
	ctx.JSON(http.StatusOK, gin.H{"status": true})
}
