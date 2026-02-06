package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := c.GetHeader("Authorization")

		// if token == "" {
		// 	c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		// 	return
		// }
		fmt.Println("Auth Middleware working")
		c.Next()
	}
}
