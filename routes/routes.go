package routes

import (
	"github.com/JammUtkarsh/cshare-server/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

var Routes = func() {
	router := SetUpRouter()
	router.POST("v1/auth/login", controller.POSTLogin)
	router.POST("v1/auth/signup", controller.POSTSignUp)
	router.PATCH("v1/users/:username/:updated", controller.UPDATEChangeUsername)
	router.PATCH("v1/users/:username", controller.UPDATEChangePassword)
	router.POST("v1/clip/:username", controller.POSTClipData)
	router.GET("v1/clip/:username/:clip_id", controller.GETClipData)
	router.GET("v1/clips/:username", controller.GETAllClipData)
	router.DELETE("v1/clips/:username/:username", controller.DELETEAllClipData)
	log.Fatalln(router.Run(":5675"))
}
