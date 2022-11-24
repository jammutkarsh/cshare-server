package routes

// routes package deals with router configuration.
// Currently, it specifically targets gin-gonic router.
// It has methods concerning ports, API version, endpoints, etc.

import (
	"log"
	"os"

	"github.com/JammUtkarsh/cshare-server/controller"
	"github.com/JammUtkarsh/cshare-server/middleware"
	"github.com/JammUtkarsh/cshare-server/utils"
	"github.com/gin-gonic/gin"
)

// SetUpRouter setting up the router Engine and returns the router.
func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// Routes defines router configuration like port and version.
var Routes = func() {
	router := SetUpRouter()
	basePath := router.Group("/v1")
	RegisterUserRoutes(basePath)
	utils.LoadEnv(".env")
	log.Fatalln(router.Run(":" + os.Getenv("SERVER_PORT")))
}

// RegisterUserRoutes defines endpoints of the server.
func RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/users")
	userRoute.POST("/signup", controller.POSTCreateUser)
	userRoute.POST("/login", controller.POSTLogin)
	userRoute.PATCH("/:username", controller.UPDATEChangePassword)

	clipRoute := rg.Group("/clip")
	// securing user routes using a Auth() function in middleware.
	secured := clipRoute.Group("/secured").Use(middleware.Auth())
	{
		secured.POST("/:username", controller.POSTClipData)
		secured.GET("/:username/:clip_id", controller.GETClipData)
		secured.GET("/:username", controller.GETAllClipData)
		secured.DELETE("/:username", controller.DELETEAllClipData)
	}

}
