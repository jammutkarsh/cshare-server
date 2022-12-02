package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JammUtkarsh/cshare-server/routes"
	"github.com/JammUtkarsh/cshare-server/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	err := log.Output(1, "error.logs")
	if err != nil {
		fmt.Println(err)
	}
	// sets router config to prod, test, debug, etc.
	utils.LoadEnv(".env")
	gin.SetMode(os.Getenv("GIN_MODE"))
}

func main() {
	routes.Routes()
}
