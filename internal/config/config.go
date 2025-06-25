package config

import (
	"strings"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/database"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/logger"
	"github.com/spf13/viper"
)

type JWTConfig struct {
	Secret               string
	Issuer               string
	AccessExpirySeconds  int64
	RefreshExpirySeconds int64
}

type Configuration struct {
	Port int
	Log  logger.LogConfig
	DB   database.DatabaseConfig
	JWT  JWTConfig
}

func GetConfig() Configuration {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// for development
	viper.SetConfigFile("app.env")
	_ = viper.ReadInConfig()

	// environment variable over .env file
	viper.AutomaticEnv()

	cfg := Configuration{
		Port: viper.GetInt("port"),
		Log: logger.LogConfig{
			Level:  viper.GetString("log_level"),
			Format: viper.GetString("log_format"),
		},
		DB: database.DatabaseConfig{
			URL: viper.GetString("database_url"),
		},
		JWT: JWTConfig{
			Secret:               viper.GetString("jwt_secret"),
			Issuer:               viper.GetString("jwt_issuer"),
			AccessExpirySeconds:  viper.GetInt64("jwt_access_expiry_seconds"),
			RefreshExpirySeconds: viper.GetInt64("jwt_refresh_expiry_seconds"),
		},
	}

	return cfg
}
