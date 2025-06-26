package database

import (
	"time"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DatabaseConfig struct {
	URL string
}

func New(cfg DatabaseConfig) *sqlx.DB {
	db, err := sqlx.Connect("mysql", cfg.URL)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	logger.Info("Database connected")
	return db
}
