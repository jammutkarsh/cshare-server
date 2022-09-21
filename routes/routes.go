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
	//TODO: Configure Middleware
	router := SetUpRouter()
	basePath := router.Group("/v1")
	RegisterUserRoutes(basePath)
	log.Fatalln(router.Run(":5675"))
}

func RegisterUserRoutes(rg *gin.RouterGroup) {
	authRoute := rg.Group("/auth")
	authRoute.POST("/login", controller.POSTLogin)
	authRoute.POST("/signup", controller.POSTCreateUser)
	userRoute := rg.Group("/users")
	userRoute.PATCH("/:username", controller.UPDATEChangePassword)
	clipRoute := rg.Group("/clip")
	clipRoute.POST("/:username", controller.POSTClipData)
	clipRoute.GET("/:username/:clip_id", controller.GETClipData)
	clipsRoute := rg.Group("/clips")
	clipsRoute.GET("/:username", controller.GETAllClipData)
	clipsRoute.DELETE("/:username", controller.DELETEAllClipData)
}
