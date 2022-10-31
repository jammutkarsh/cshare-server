package routes

import (
	"github.com/JammUtkarsh/cshare-server/controller"
	"github.com/JammUtkarsh/cshare-server/middleware"
	"github.com/JammUtkarsh/cshare-server/utils"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

var Routes = func() {
	router := SetUpRouter()
	basePath := router.Group("/v1")
	RegisterUserRoutes(basePath)
	utils.LoadEnv(".env")
	log.Fatalln(router.Run(":" + os.Getenv("SERVER_PORT")))
}

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/users")
	userRoute.POST("/signup", controller.POSTCreateUser)
	userRoute.POST("/login", controller.POSTLogin)
	userRoute.PATCH("/:username", controller.UPDATEChangePassword)

	clipRoute := rg.Group("/clip")
	secured := clipRoute.Group("/secured").Use(middleware.Auth())
	{
		secured.POST("/:username", controller.POSTClipData)
		secured.GET("/:username/:clip_id", controller.GETClipData)
		secured.GET("/:username", controller.GETAllClipData)
		secured.DELETE("/:username", controller.DELETEAllClipData)
	}

}
