package http

import (
	"github.com/gin-gonic/gin"
)

// Users represents a type that can provide CRUD operations on users.
type Users interface {
	GetUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	GetUserByID(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type Server struct {
	users Users
}

// New will instantiate a new instance of Server.
func New(u Users) *Server {
	return &Server{
		users: u,
	}
}

// AddRoutes will add the routes this server supports to the router.
func (s *Server) AddRoutes(r *gin.Engine) {
	grp := r.Group("/user-api")
	{
		grp.GET("user", s.users.GetUsers)
		grp.POST("user", s.users.CreateUser)
		grp.GET("user/:id", s.users.GetUserByID)
		grp.PUT("user/:id", s.users.UpdateUser)
		grp.DELETE("user/:id", s.users.DeleteUser)
	}
}
