package main

import (
	"fmt"
	"github.com/JammUtkarsh/cshare-server/routes"
	"log"
)

func init() {
	err := log.Output(1, "error.logs")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	routes.Routes()
}
