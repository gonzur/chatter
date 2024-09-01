package main

import (
	"chatter-server/internal/auth"
	"chatter-server/internal/chatrooms"
	"log"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "http://localhost")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	serv := gin.Default()
	serv.Use(CORSMiddleware())

	apiBaseRoute := serv.Group("/api")
	apiBaseRoute.POST("/login", auth.Login)
	apiBaseRoute.POST("/create-user", auth.CreateUser)
	chatrooms.AttachRoutes(apiBaseRoute)

	if err := serv.Run(":8080"); err != nil {
		log.Println(err.Error())
	}
}
