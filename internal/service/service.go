package service

import "github.com/gin-gonic/gin"

// User represents a user in the system.
type User struct {
	ID       int
	Username string
}

// UserService defines the interface for user-related operations.
type UserService interface {
	GetUserByID(id string) User
	GetUserByUsername(username string) User
}

// BaseUserService is a basic implementation of the UserService interface.
// in real applications, this would interact with a database.
type BaseUserService struct {
	UserData map[string]User
}

// GetUserByID retrieves a user by their ID.
func (s *BaseUserService) GetUserByID(id string) User {
	return s.UserData[id]
}

// GetUserByUsername retrieves a user by their username.
func (s *BaseUserService) GetUserByUsername(username string) User {
	for _, user := range s.UserData {
		if user.Username == username {
			return user
		}
	}
	return User{}
}

// UserServiceHandler handles HTTP requests related to users.
type UserServiceHandler struct {
	Service UserService
}

// FetchUserByID handles a request to fetch a user by their ID.
func (h *UserServiceHandler) FetchUserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(
		200,
		h.Service.GetUserByID(id),
	)
}

// FetchUserByUsername handles a request to fetch a user by their username.
func (h *UserServiceHandler) FetchUserByUsername(c *gin.Context) {
	username := c.Param("username")
	c.JSON(
		200,
		h.Service.GetUserByUsername(username),
	)
}
