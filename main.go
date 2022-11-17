package main

import (
	"fmt"
	"github.com/JammUtkarsh/cshare-server/routes"
	"github.com/JammUtkarsh/cshare-server/utils"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func init() {
	err := log.Output(1, "error.logs")
	if err != nil {
		fmt.Println(err)
	}
	utils.LoadEnv(".env")
	gin.SetMode(os.Getenv("GIN_MODE"))
}

func main() {
	routes.Routes()
}
