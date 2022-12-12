package main

import (
	"chatter-server/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	serv := gin.Default()
	routes.BuildRoutes(serv)

	if err := serv.Run(":8080"); err != nil {
		println(err.Error())
	}
}
