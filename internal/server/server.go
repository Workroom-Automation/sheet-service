package server

import (
	"github.com/gin-gonic/gin"
	"github.com/leapsquare/sheet-service/config"
	"github.com/leapsquare/sheet-service/pkg/logger"
	"gorm.io/gorm"

	"os"
	"os/signal"
	"syscall"
)

const (
	ctxTimeout = 5
)

// Server struct
type Server struct {
	cfg      *config.Config
	postgres *gorm.DB
	logger   logger.Logger
	router   *gin.Engine
}

// NewServer New Server constructor
func NewServer(cfg *config.Config, logger logger.Logger, postgres *gorm.DB) *Server {
	router := gin.Default()
	return &Server{
		router:   router,
		cfg:      cfg,
		logger:   logger,
		postgres: postgres,
	}
}

func (s *Server) Run() error {
	// TODO add auth
	//tokenVerifier := auth.NewTokenVerifierMiddleware(s.cfg.Auth.Auth0KeySetUrl)
	//s.app.Use(tokenVerifier.Verify)
	if err := s.MapHandlers(); err != nil {
		return err
	}
	go func() {
		//if err := s.app.Listen(s.cfg.Server.Port); err != nil {
		//	s.logger.Errorf("Error listening to Port: %s", err)
		//}
		if err := s.router.RunTLS(s.cfg.Server.Port, "localhost.crt", "localhost.key"); err != nil {
			s.logger.Errorf("Error listening to Port: %s", err)
		}

	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	//ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	//defer shutdown()
	s.logger.Info("Server Exited Properly")
	//return s.app.Shutdown(ctx)
	return nil
}
