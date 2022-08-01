package controller

import (
	"fmt"
	"github.com/JammUtkarsh/cshare-server/models"
	"github.com/JammUtkarsh/cshare-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	formatValidationErrType = "format_validation_error"
	credValidationErrType   = "credential_validation_error"
	resourceNotFoundErrType = "resource_not_found_error"
	userNotFoundErrType     = "user_not_found_error"
	serviceErrType          = "service_error"
)

// TODO for all fmt.Errorf()
// log error in a file or different format.
// Avoid mixing these log with middleware logs.

// TODO: for all 200 response codes, provide a better response to the client.

// TODO: Use cont.body(send data in json format) to get the data from client and vise versa.

// TODO: add headers to the response which include some auth token or other essential information.
//<--User Authentication and Authorisation-->

// POSTLogin gets the user data from client and inserts it into the database for the first time.
func POSTLogin(cont *gin.Context) {
	var clientCredentials models.Users
	if err := cont.BindJSON(&clientCredentials); err != nil {
		cont.JSON(http.StatusBadRequest, gin.H{"error": formatValidationErrType})
		utils.ErrorReaderWriter(err, POSTLogin)
	} else {
		err, val := Authenticate(clientCredentials.Username, clientCredentials.Password)
		if val == false {
			utils.ErrorReaderWriter(err, POSTLogin)
			cont.JSON(http.StatusUnauthorized, gin.H{"status": credValidationErrType})
		} else {
			cont.JSON(http.StatusOK, gin.H{
				"status":     val,
				"username":   clientCredentials.Username,
				"token_type": "Bearer",
				"token":      "",
			})
		}
	}
}

// POSTSignUp gets the user data from client and inserts it into the database for the first time.
func POSTSignUp(cont *gin.Context) {
	db := models.CreateConnection()
	var clientCredentials models.Users
	// password is hashed and then marshaled to json before sending to the client.
	if err := cont.BindJSON(&clientCredentials); err != nil {
		cont.JSON(http.StatusBadRequest, gin.H{"error": formatValidationErrType})
		utils.ErrorReaderWriter(err, POSTSignUp)
	} else {
		err, _ := models.InsertUser(db, clientCredentials.Username)
		err, val := Authorize(clientCredentials.Username, clientCredentials.Password)
		if val == false {
			cont.JSON(http.StatusUnauthorized, gin.H{"error": userNotFoundErrType})
			utils.ErrorReaderWriter(err, POSTSignUp)
		} else {
			cont.JSON(http.StatusOK, gin.H{"status": val, "username": clientCredentials.Username})
		}
	}
	models.CloseConnection(db)
}

func UPDATEChangeUsername(cont *gin.Context) {
	db := models.CreateConnection()
	initialName := cont.Param("username")
	finalName := cont.Param("updated")
	if err, val := models.SelectByUsername(db, finalName); val != -1 {
		cont.JSON(http.StatusBadRequest, gin.H{"error": userNotFoundErrType})
		utils.ErrorReaderWriter(err, UPDATEChangeUsername)
	} else {
		err, _ := models.UpdateByUsername(db, initialName, finalName)
		utils.ErrorReaderWriter(err, UPDATEChangeUsername)
		cont.JSON(http.StatusOK, gin.H{"status": val})
	}
	models.CloseConnection(db)
}

func UPDATEChangePassword(cont *gin.Context) {
	db := models.CreateConnection()

	models.CloseConnection(db)
}

//<--End-->

//<--Data related Functions-->

// POSTClipData gets the clip data from client and inserts it into the database.
// Client is required to send the user_id, message, and secret as a JSON object.
func POSTClipData(cont *gin.Context) {
	var clientData models.ClipStack
	db := models.CreateConnection()
	if err := cont.BindJSON(&clientData); err != nil {
		cont.JSON(http.StatusBadRequest, gin.H{"error": formatValidationErrType})
		_ = fmt.Errorf("1  %v", err)
	} else {
		err, val := models.InsertClip(db, clientData)
		utils.ErrorReaderWriter(err, POSTClipData)
		if err != nil {
			utils.ErrorReaderWriter(err, POSTClipData)
			cont.JSON(http.StatusBadRequest, gin.H{"error": err})
		} else {
			cont.JSON(http.StatusOK, gin.H{"status": val})
		}
	}
	models.CloseConnection(db)
}

func GETClipData(cont *gin.Context) {
	db := models.CreateConnection()
	uname := cont.Param("username")
	if err, val := models.SelectByUsername(db, uname); val == -1 {
		cont.JSON(http.StatusUnauthorized, gin.H{"error": userNotFoundErrType})
		utils.ErrorReaderWriter(err, GETClipData)
	} else {
		clipId := cont.Param("clip_id")
		clipIdInt, _ := strconv.ParseInt(clipId, 10, 64)
		err, count := models.ClipCount(db, val)
		utils.ErrorReaderWriter(err, GETClipData)
		if count >= clipIdInt {
			err, data := models.SelectClip(db, clipIdInt, val)
			if err != nil {
				return
			}
			cont.SecureJSON(http.StatusOK, data)
		} else {
			cont.JSON(http.StatusBadRequest, gin.H{"error": resourceNotFoundErrType})
		}
	}
	models.CloseConnection(db)
}

func GETAllClipData(cont *gin.Context) {
	db := models.CreateConnection()
	var dataSet []models.ClipStack
	uname := cont.Param("username")
	if err, val := models.SelectByUsername(db, uname); val == -1 {
		cont.JSON(http.StatusUnauthorized, gin.H{"error": userNotFoundErrType})
		if err != nil {
			utils.ErrorReaderWriter(err, GETAllClipData)
		}
	} else {
		err, count := models.ClipCount(db, val)
		if err != nil {
			utils.ErrorReaderWriter(err, GETAllClipData)
		}
		for i := 0; int64(i) < count; i++ {
			err, data := models.SelectClip(db, count-1, val)
			utils.ErrorReaderWriter(err, GETAllClipData)
			count--
			dataSet = append(dataSet, data)
		}
		cont.JSON(http.StatusOK, dataSet)
	}
	models.CloseConnection(db)
}

func DELETEAllClipData(cont *gin.Context) {
	db := models.CreateConnection()
	uname := cont.Param("username")
	if err, val := models.SelectByUsername(db, uname); val == -1 {
		cont.JSON(http.StatusUnauthorized, gin.H{"error": "user doesn't exists"})
		utils.ErrorReaderWriter(err, DELETEAllClipData)
	} else {
		var i int64
		err, count := models.ClipCount(db, val)
		utils.ErrorReaderWriter(err, DELETEAllClipData)
		for i = 0; i <= count; i++ {
			err := models.DeleteClip(db, int64(i), val)
			utils.ErrorReaderWriter(err, DELETEAllClipData)
		}
		cont.JSON(http.StatusOK, gin.H{"status": true})
	}
	models.CloseConnection(db)
}

//<--End-->
