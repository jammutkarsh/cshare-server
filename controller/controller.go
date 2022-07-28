package controller

import (
	"fmt"
	"github.com/JammUtkarsh/cshare-server/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Set constants of error messages to return to the client.
const (
	errUser = "usernotfound"
)

// TODO for all fmt.Errorf()
// log error in a file or different format.
// Avoid mixing these log with middleware logs.

// TODO: for all 200 response codes, provide a better response to the client.

//<--User Authentication and Authorisation-->

// POSTLogin gets the user data from client and inserts it into the database for the first time.
func POSTLogin(cont *gin.Context) {
	var clientCredentials models.Users
	if err := cont.BindJSON(&clientCredentials); err != nil {
		cont.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
		_ = fmt.Errorf("%v", err)
	} else {
		err2, val := Authorize(clientCredentials.Username, clientCredentials.Password)
		if val == false {
			_ = fmt.Errorf("%v", err2)
			cont.JSON(http.StatusUnauthorized, gin.H{"status": val})
		} else {
			cont.JSON(http.StatusOK, gin.H{"status": val})
		}
	}
}

// POSTSignUp gets the user data from client and inserts it into the database for the first time.
func POSTSignUp(cont *gin.Context) {
	db := models.CreateConnection()
	var clientCredentials models.Users
	if err := cont.BindJSON(&clientCredentials); err != nil {
		cont.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
		_ = fmt.Errorf("%v", err)
	} else {
		err, _ := models.InsertUser(db, clientCredentials.Username)
		err, val := Authenticate(clientCredentials.Username, clientCredentials.Password)
		if val == false {
			cont.JSON(http.StatusUnauthorized, gin.H{"error": "uname exists"})
			_ = fmt.Errorf("%v", err)
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
		cont.JSON(http.StatusBadRequest, gin.H{"error": "username exists"})
		_ = fmt.Errorf("%v", err)
	} else {
		err, _ := models.UpdateByUsername(db, initialName, finalName)
		_ = fmt.Errorf("%v", err)
		cont.JSON(http.StatusOK, gin.H{"status": true})
	}
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
		cont.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!1"})
		_ = fmt.Errorf("1  %v", err)
	} else {
		err, val := models.InsertClip(db, clientData)
		_ = fmt.Errorf("2  %v", err)
		if err != nil {
			_ = fmt.Errorf("%v", err)
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
		cont.JSON(http.StatusUnauthorized, gin.H{"error": "username exists"})
		_ = fmt.Errorf("%v", err)
	} else {
		clipId := cont.Param("clip_id")
		clipIdInt, _ := strconv.ParseInt(clipId, 10, 64)
		err, count := models.ClipCount(db, val)
		_ = fmt.Errorf("%v", err)
		if count >= clipIdInt {
			err, data := models.SelectClip(db, clipIdInt, val)
			if err != nil {
				return
			}
			cont.JSON(http.StatusOK, data)
		} else {
			cont.JSON(http.StatusBadRequest, gin.H{"error": "clip_id not found"})
		}
	}
	models.CloseConnection(db)
}

func GETAllClipData(cont *gin.Context) {
	db := models.CreateConnection()
	var dataSet []models.ClipStack
	uname := cont.Param("username")
	err, id := models.GetUserID(db, uname)
	if err != nil {
		_ = fmt.Errorf("%v", err)

	}
	err, count := models.ClipCount(db, id)
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}
	for i := 0; int64(i) < count; i++ {
		err, data := models.SelectClip(db, count-1, id)
		_ = fmt.Errorf("%v", err)
		count--
		dataSet = append(dataSet, data)
	}
	//byteData, err := json.MarshalIndent(data, "", " ")
	cont.JSON(http.StatusOK, dataSet)
	models.CloseConnection(db)
}

func DELETEClipData(cont *gin.Context) {
	db := models.CreateConnection()
	uname := cont.Param("username")
	if err, val := models.SelectByUsername(db, uname); val == -1 {
		cont.JSON(http.StatusUnauthorized, gin.H{"error": "user doesn't exists"})
		_ = fmt.Errorf("%v", err)
	} else {
		clipId := cont.Param("clip_id")
		clipIdInt, _ := strconv.ParseInt(clipId, 10, 64)
		err := models.DeleteClip(db, clipIdInt, val)
		_ = fmt.Errorf("%v", err)
		cont.JSON(http.StatusOK, gin.H{"status": true})
	}
	models.CloseConnection(db)
}
func DELETEAllClipData(cont *gin.Context) {
	db := models.CreateConnection()
	uname := cont.Param("username")
	if err, val := models.SelectByUsername(db, uname); val == -1 {
		cont.JSON(http.StatusUnauthorized, gin.H{"error": "user doesn't exists"})
		_ = fmt.Errorf("%v", err)
	} else {
		var i int64
		err, count := models.ClipCount(db, val)
		_ = fmt.Errorf("%v", err)
		for i = 0; i <= count; i++ {
			err := models.DeleteClip(db, int64(i), val)
			_ = fmt.Errorf("%v", err)
		}
		cont.JSON(http.StatusOK, gin.H{"status": true})
	}
	models.CloseConnection(db)
}

//<--End-->
