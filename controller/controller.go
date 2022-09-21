package controller

import (
	"errors"
	"github.com/JammUtkarsh/cshare-server/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

const (
	formatValidationErrType = "format_validation_error"
	credValidationErrType   = "credential_validation_error"
	resourceNotFoundErrType = "resource_not_found_error"
	userNotFoundErrType     = "user_not_found_error"
	authErrType             = "authentication_error"
	serviceErrType          = "service_error"
)

// TODO: for all 200 response codes, provide a better response to the client.

// TODO: add headers to the response which include some auth token or other essential information.
//<--User Authentication and Authorisation-->

// POSTLogin gets the user data from client and inserts it into the database for the first time.
func POSTLogin(ctx *gin.Context) {
	var clientCredentials models.Users
	if err := ctx.BindJSON(&clientCredentials); err != nil { // checking if the json response is in the correct format.
		ctx.JSON(http.StatusBadRequest, gin.H{"error": formatValidationErrType})
		log.Println(err)
	} else { // format is correct ; now checking if the credentials are correct.
		err, val := Authenticate(clientCredentials.Username, clientCredentials.Password)
		if val == false {
			log.Println(err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": credValidationErrType})
		} else { // credentials are correct ; giving status:200 with auth token.
			ctx.JSON(http.StatusOK, gin.H{
				"status":     val,
				"username":   clientCredentials.Username,
				"token_type": "Bearer",
				"token":      "",
			})
		}
	}
}

// POSTCreateUser gets the user data from client and inserts it into the database for the first time.
// TODO: rewrite the function after creating authentication system.
func POSTCreateUser(ctx *gin.Context) {
	db := models.CreateConnection()
	models.CloseConnection(db)
	var clientCredentials models.Users
	if err := ctx.BindJSON(&clientCredentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": formatValidationErrType})
		log.Println(err)
	} else {
		err, _ := models.InsertUser(db, clientCredentials.Username)
		log.Println(err)
		err, val := Authorize(clientCredentials.Username, clientCredentials.Password)
		if val == false {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": authErrType})
			log.Println(err)
		} else {
			ctx.JSON(http.StatusOK, gin.H{"status": val, "username": clientCredentials.Username})
		}
	}
}

func UPDATEChangePassword(ctx *gin.Context) {
	db := models.CreateConnection()
	models.CloseConnection(db)

}

//<--End-->

//<--Data related Functions-->

// POSTClipData gets the clip data from client and inserts it into the database.
// Client is required to send the user_id, message, and secret as a JSON object.
func POSTClipData(ctx *gin.Context) {
	var clientData models.Data
	db := models.CreateConnection()
	models.CloseConnection(db)
	if err := ctx.BindJSON(&clientData); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New(formatValidationErrType))
		log.Println(err)
		return
	}
	clip := models.Data{
		Username: clientData.Username,
		Message:  clientData.Message,
		Secret:   clientData.Secret,
	}
	if err, _ := models.InsertClip(db, clip); err != nil {
		_ = ctx.AbortWithError(http.StatusNotFound, errors.New(serviceErrType))
		return
	}
	ctx.JSON(http.StatusCreated, &clientData)

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
	err, val := models.SelectByUsername(db, ctx.Param("username"))
	if err != nil {
		log.Println(err)
	}
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
	err, val := models.SelectByUsername(db, ctx.Param("username"))
	if val == -1 {
		_ = ctx.AbortWithError(http.StatusUnauthorized, errors.New(userNotFoundErrType))
		log.Println(err)
	}
	_, count := models.ClipCount(db, val)
	log.Println(err)
	for i := int64(0); i <= count; i++ {
		err := models.DeleteClip(db, i, val)
		if err != nil {
			log.Println(err)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true})
}

//<--End-->
