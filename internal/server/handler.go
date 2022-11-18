package server

import (
	"github.com/gin-gonic/gin"
	"github.com/leapsquare/sheet-service/internal/middleware"

	"net/http"
)

func (s *Server) MapHandlers() error {
	mdlwr := middleware.NewMiddleware(s.cfg, s.logger)

	// Init repo

	// Init service

	// Init handlers

	// Enabling routes
	v1 := s.router.Group("/api/v1")

	health := s.router.Group("/sheet-backend/health")
	health.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})
	return nil
}
