package main

import (
	"chatter-server/internal/chatrooms"
	"chatter-server/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	serv := gin.Default()
	routes.BuildRoutes(serv)
	if serv.Run() != nil {
		return
	}
	room := chatrooms.Room{}
	room.Init("")
}
