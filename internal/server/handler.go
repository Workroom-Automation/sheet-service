package server

import (
	"github.com/gin-gonic/gin"
	"github.com/leapsquare/sheet-service/internal/middleware"
	"github.com/leapsquare/sheet-service/internal/sheet"
	"net/http"
)

func (s *Server) MapHandlers() error {
	mdlwr := middleware.NewMiddleware(s.cfg, s.logger)

	// Init repo
	sheetRepo := sheet.NewRepository(s.logger)
	// Init service
	sheetSvc := sheet.NewService(s.logger, s.postgres, s.cfg, sheetRepo)
	// Init handlers
	sheetHandler := sheet.NewController(s.cfg, s.logger, sheetSvc)
	// Enabling routes
	v1 := s.router.Group("/api/v1")
	sheet.MapRoutesV1(v1, sheetHandler, mdlwr)

	health := s.router.Group("/sheet-backend/health")
	health.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})
	return nil
}
