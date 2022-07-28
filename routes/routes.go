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
	router.POST("/login", controller.POSTLogin)
	router.POST("/signup", controller.POSTSignUp)
	router.POST("/postclip", controller.POSTClipData)
	router.GET("/getclip/:username", controller.GETAllClipData)
	router.GET("/getclip/:username/:clip_id", controller.GETClipData)
	router.PUT("/updateusername/:username/:updated", controller.UPDATEChangeUsername)
	router.DELETE("/deleteclip/:username/:clip_id", controller.DELETEClipData)
	router.DELETE("/deleteall/:username", controller.DELETEAllClipData)
	log.Fatalln(router.Run(":5675"))
}
