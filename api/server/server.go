// Package server defined the server main operation including
// creating dependencies which are related to the server.
package server

import (
	"fmt"

	"github.com/KryptoKnight/backend-test-golang/api/server/handlers"
	"github.com/KryptoKnight/backend-test-golang/api/server/middleware"
	"github.com/KryptoKnight/backend-test-golang/config"
	"github.com/KryptoKnight/backend-test-golang/repo"
	"github.com/KryptoKnight/backend-test-golang/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// Server contains server specific details.
type Server struct {
	addr   string
	engine *gin.Engine
}

// NewServer creates a new instance of server.
func NewServer() *Server {
	return &Server{
		engine: gin.Default(),
	}
}

// Configure configure server.
func (s *Server) Configure(cfg *config.App, logger zerolog.Logger) {
	serCfg := cfg.Server

	// Server configurations
	s.addr = fmt.Sprintf(":%v", serCfg.Port)

	// Configure repository
	repo := repo.NewFactMongoRepo()

	// Configure service
	factService := service.NewFactService(repo, logger)

	// Configure handlers
	rootRouter := s.engine.Group("/api")
	rootRouter.Use(middleware.GetJwtValidationMiddleware(serCfg.PublicKeyPath, logger))

	// v1 routes
	v1Router := rootRouter.Group("/v1")
	v1Router.GET("/facts", handlers.GetRetrieveFactsHandler(factService))

	// v2 routes
	v2Router := rootRouter.Group("/v2")
	v2Router.GET("/facts", handlers.GetNotImplementedHandler())

}

// Start server
func (s *Server) Start() {
	go func() {
		s.engine.Run(s.addr)
	}()
}
