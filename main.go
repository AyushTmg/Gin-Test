package main

import (
	"Gin/internal/handler"
	"Gin/internal/service"
	"Gin/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/", handler.Home)

	// static data
	staticData := map[string]service.User{
		"1": {ID: 1, Username: "Alice Smith"},
		"2": {ID: 2, Username: "Bob Jones"},
		"3": {ID: 3, Username: "Charlie Brown"},
	}

	// initialize the Base UserService with static data to simulate database interactions
	userService := &service.BaseUserService{UserData: staticData}
	userServiceHandler := &service.UserServiceHandler{Service: userService}

	// define routes for fetching users by ID and username
	/*
		! Note:
		the routes cannot be like user/:id and user/:username because they
		will conflict with each other. So we need to differentiate them by adding
		a prefix like in this case

	*/
	app.GET("users/id/:id", userServiceHandler.FetchUserByID)
	app.GET("users/username/:username", userServiceHandler.FetchUserByUsername)

	auth := app.Group("/auth", middleware.AuthMiddleware())
	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)

	app.Run(":8080")
}
