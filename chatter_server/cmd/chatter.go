package main

import (
	"chatter-server/internal/chatrooms"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	serv := gin.Default()

	apiRouter := serv.Group("/api")
	chatrooms.AttachRoutes(apiRouter)

	if err := serv.Run(":8080"); err != nil {
		log.Println(err.Error())
	}
}
