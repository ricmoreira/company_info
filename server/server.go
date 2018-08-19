package server

import (
	"company_info/config"
	"company_info/controllers/v1"
	"company_info/handlers"

	"github.com/gin-gonic/gin"
)

// Server is the http layer for role and user resource
type Server struct {
	config           *config.Config
	headerController *controllers.HeaderController
	handlers         *handlers.HttpHandlers
}

// NewServer is the Server constructor
func NewServer(cf *config.Config,
	hc *controllers.HeaderController,
	hand *handlers.HttpHandlers) *Server {

	return &Server{
		config:           cf,
		headerController: hc,
		handlers:         hand,
	}
}

// Run loads server with its routes and starts the server
func (s *Server) Run() {
	// Instantiate a new router
	r := gin.Default()

	// generic routes
	r.HandleMethodNotAllowed = false
	r.NoRoute(s.handlers.NotFound)

	// Header resource
	headerApi := r.Group("/api/v1/header")
	{
		// Create a new header
		headerApi.POST("", s.headerController.CreateAction)

		// List headers with filtering and pagination
		headerApi.GET("", s.headerController.ListAction)
	}

	// Fire up the server
	r.Run(s.config.Host)
}
