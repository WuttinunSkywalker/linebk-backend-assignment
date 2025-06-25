package logger

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LogConfig struct {
	Level  string
	Format string
}

func Init(cfg LogConfig) {
	logLevelStr := strings.ToLower(cfg.Level)
	logLevel, err := zerolog.ParseLevel(logLevelStr)
	if err != nil || logLevelStr == "" {
		logLevel = zerolog.InfoLevel
	}

	logFormat := strings.ToLower(cfg.Format)

	if logFormat == "json" {
		// For production, use fast JSON logging.
		log.Logger = zerolog.New(os.Stdout).
			Level(logLevel).
			With().
			Timestamp().
			Logger()
	} else {
		// For development, use console logging.
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		log.Logger = zerolog.New(output).
			Level(logLevel).
			With().
			Timestamp().
			Logger()
	}

	log.Info().Msgf("Logger initialized with level '%s'", logLevel)
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Infof(format string, v ...any) {
	log.Info().Msgf(format, v...)
}

func Debug(msg string) {
	log.Debug().Msg(msg)
}

func Debugf(format string, v ...any) {
	log.Debug().Msgf(format, v...)
}

func Warn(msg string) {
	log.Warn().Msg(msg)
}

func Warnf(format string, v ...any) {
	log.Warn().Msgf(format, v...)
}

func Error(v any) {
	switch val := v.(type) {
	case error:
		log.Error().Msg(val.Error())
	case string:
		log.Error().Msg(val)
	default:
		log.Error().Msgf("%v", val)
	}
}

func Errorf(format string, v ...any) {
	log.Error().Msgf(format, v...)
}

func Fatal(msg string) {
	log.Fatal().Msg(msg)
}

func Fatalf(format string, v ...any) {
	log.Fatal().Msgf(format, v...)
}