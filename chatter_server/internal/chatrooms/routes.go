package chatrooms

import (
	"github.com/gin-gonic/gin"
)

func AttachRoutes(router *gin.RouterGroup) {
	roomSetup := GinRoute

	chatRouter := router.Group("/chat")
	chatRouter.GET("/join-room", roomSetup)
}
