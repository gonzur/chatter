package main

import (
	"chatter-server/internal/chatrooms"

	"github.com/gin-gonic/gin"
)

func main() {
	serv := gin.Default()

	apiRouter := serv.Group("/api")
	chatrooms.AttachRoutes(apiRouter)

	if err := serv.Run(":8080"); err != nil {
		println(err.Error())
	}
}
