package main

import (
	"fmt"
	"github.com/JammUtkarsh/cshare-server/routes"
	"log"
)

func main() {
	err := log.Output(1, "error.logs")
	if err != nil {
		fmt.Println(err)
	}
	routes.Routes()
}
