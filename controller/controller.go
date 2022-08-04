package controller

import (
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
func POSTLogin(cont *gin.Context) {
	var clientCredentials models.Users
	if err := cont.BindJSON(&clientCredentials); err != nil { // checking if the json response is in the correct format.
		cont.JSON(http.StatusBadRequest, gin.H{"error": formatValidationErrType})
		log.Println(err)
	} else { // format is correct ; now checking if the credentials are correct.
		err, val := Authenticate(clientCredentials.Username, clientCredentials.Password)
		if val == false {
			log.Println(err)
			cont.JSON(http.StatusUnauthorized, gin.H{"status": credValidationErrType})
		} else { // credentials are correct ; giving status:200 with auth token.
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
// TODO: rewrite the function after creating authentication system.
func POSTSignUp(cont *gin.Context) {
	db := models.CreateConnection()
	var clientCredentials models.Users
	if err := cont.BindJSON(&clientCredentials); err != nil {
		cont.JSON(http.StatusBadRequest, gin.H{"error": formatValidationErrType})
		log.Println(err)
	} else {
		err, _ := models.InsertUser(db, clientCredentials.Username)
		log.Println(err)
		err, val := Authorize(clientCredentials.Username, clientCredentials.Password)
		if val == false {
			cont.JSON(http.StatusUnauthorized, gin.H{"error": authErrType})
			log.Println(err)
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
		log.Println(err)
	} else {
		err, _ := models.UpdateByUsername(db, initialName, finalName)
		log.Println(err)
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
		log.Println(err)
	} else {
		username := cont.Param("username")
		err, val := models.InsertClip(db, clientData, username)
		if err != nil {
			log.Println(err)
			cont.JSON(http.StatusBadRequest, gin.H{"error": err})
		} else {
			cont.JSON(http.StatusOK, gin.H{"status": val,
				"data": clientData})
		}
	}
	models.CloseConnection(db)
}

func GETClipData(cont *gin.Context) {
	db := models.CreateConnection()
	uname := cont.Param("username")
	if err, val := models.SelectByUsername(db, uname); val == -1 {
		cont.JSON(http.StatusUnauthorized, gin.H{"error": userNotFoundErrType})
		log.Println(err)
	} else {
		clipId := cont.Param("clip_id")
		clipIdInt, _ := strconv.ParseInt(clipId, 10, 64)
		err, count := models.ClipCount(db, val)
		log.Println(err)
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
			log.Println(err)
		}
	} else {
		err, count := models.ClipCount(db, val)
		if err != nil {
			log.Println(err)
		}
		for i := 0; int64(i) < count; i++ {
			err, data := models.SelectClip(db, count-1, val)
			log.Println(err)
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
		cont.JSON(http.StatusUnauthorized, gin.H{"error": userNotFoundErrType})
		log.Println(err)
	} else {
		var i int64
		err, count := models.ClipCount(db, val)
		log.Println(err)
		for i = 0; i <= count; i++ {
			err := models.DeleteClip(db, i, val)
			log.Println(err)
		}
		cont.JSON(http.StatusOK, gin.H{"status": true})
	}
	models.CloseConnection(db)
}

//<--End-->
