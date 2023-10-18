package http

import (
	"github.com/gin-gonic/gin"
)

// Service represents a http service that provides routes for the listener.
type Service interface {
	AddRoutes(r *gin.Engine)
}

// New instantiates a new instance of Server.
func New(s Service) error {
	r := gin.Default()

	s.AddRoutes(r)

	return r.Run()
}
