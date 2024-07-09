package chatrooms

import (
	"github.com/gin-gonic/gin"
)

func AttachRoutes(router *gin.RouterGroup) {
	chatRouter := router.Group("/chat")

	chatRouter.GET("/join", RoomSetup)
	chatRouter.GET("/list", ActiveRooms)
}
