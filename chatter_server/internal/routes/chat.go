package routes

import "github.com/gin-gonic/gin"

func chatRoutes(router *gin.RouterGroup) {
	chatRouter := router.Group("/chat")
	chatRouter.GET("/message")

}

func BuildRoutes(router *gin.Engine) {
	apiRouter := router.Group("/api")
	chatRoutes(apiRouter)
}
