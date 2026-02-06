package main

import (
	"Gin/internal/handler"
	"Gin/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/", handler.Home)

	auth := app.Group("/auth", middleware.AuthMiddleware())
	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)

	app.Run(":8080")
}
