package routes

import (
	"chatter-server/internal/chatrooms"
	"net/http"

	"github.com/gin-gonic/gin"
)

func joinRoom(router *gin.RouterGroup) {
	chatrooms.Init()
	setup := func(c *gin.Context) {

		conn, err := chatrooms.Upgrade(c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		if err = chatrooms.Create(""); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		if err = chatrooms.Join("", "", conn); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

	}

	chatRouter := router.Group("/chat")
	chatRouter.GET("/joinRoom", setup)
}

func BuildRoutes(router *gin.Engine) {
	apiRouter := router.Group("/api")
	joinRoom(apiRouter)
}
