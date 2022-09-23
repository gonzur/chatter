package routes

import (
	"chatter-server/internal/chatrooms"

	"github.com/gin-gonic/gin"
)

func joinRoom(router *gin.RouterGroup) {
	chatrooms.Init()
	roomSetup := chatrooms.GinRoute

	chatRouter := router.Group("/chat")
	chatRouter.GET("/joinRoom", roomSetup)
}

func BuildRoutes(router *gin.Engine) {
	apiRouter := router.Group("/api")
	joinRoom(apiRouter)
}
