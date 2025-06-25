package database

import (
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

	err = db.Ping()
	if err != nil {
		db.Close()
		logger.Fatalf("Failed to ping database: %v", err)
	}

	logger.Info("Database connected")
	return db
}
