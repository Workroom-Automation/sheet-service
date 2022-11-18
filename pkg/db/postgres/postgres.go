package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConnectionConfig struct {
	DSN                string
	MaxIdleConnections int
	MaxOpenConnections int
}

func InitPostgres(config *PostgresConnectionConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("unable to establish connection with postgres: %w ", err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("unable to establish connection with postgres: %w ", err))
	}
	if config.MaxIdleConnections > 0 {
		sqlDB.SetMaxIdleConns(config.MaxIdleConnections)
	}
	if config.MaxOpenConnections > 0 {
		sqlDB.SetMaxOpenConns(config.MaxOpenConnections)
	}
	return db
}
