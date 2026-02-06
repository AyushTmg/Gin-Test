package handler

import (
	"Gin/internal/dto"

	"github.com/gin-gonic/gin"
)

// ! Home Page Handler
func Home(c *gin.Context) {
	/*
		 Everything revolves around *gin.Context.
			c.Param("id")
			c.Query("page")
			c.BindJSON(&dto)
			c.JSON(200, data)
	*/
	c.JSON(200, gin.H{
		"message": "Welcome to HomePage",
	})
}

// ! User Registration Handler
func Register(c *gin.Context) {
	// This Allocate a struct in memory
	var req dto.CreateUserRequest

	/*
		c.ShouldBindJSON(&req) this reads request body and try to bind it to the struct,
		if it fails it returns an error

		And also notice that we have used &req because we want to pass the reference of the struct
		to the function so that it can modify the struct and fill it with the data from the request body.
	*/
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":   "Register",
		"full_name": req.FullName,
		"email":     req.Email,
	})

	/*
		This will return the whole struct as a response which we don't
		want to do in real application because it contains password field
	*/
	// c.JSON(200, req)

}

func Login(c *gin.Context) {
	var req dto.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Login successful",
		"email":   req.Email,
	})
}
