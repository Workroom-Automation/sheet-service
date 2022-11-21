package api

import (
	"github.com/leapsquare/sheet-service/internal/server"
	"github.com/leapsquare/sheet-service/pkg/db/postgres"
	"github.com/leapsquare/sheet-service/pkg/logger"
	"github.com/leapsquare/sheet-service/pkg/utils"
	"log"
)

// Init @title Sheet Service Backend Documentation
// @contact.name Rohan Agarwal
// @BasePath /api
func Init(args []string) error {
	log.Println("Starting api server")
	cfg, err := utils.LoadAndParseCfgFile()
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	db := postgres.InitPostgres(&postgres.PostgresConnectionConfig{
		DSN: cfg.Postgres.DSN,
	})
	if err = server.NewServer(cfg, appLogger, db).Run(); err != nil {
		appLogger.Fatalf("error running the server, err: %v", err)
	}
	return nil
}
