package logger_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	testCases := []struct {
		name           string
		config         logger.LogConfig
		expectedLevel  zerolog.Level
		expectedFormat string
	}{
		{
			name: "JSONFormatInfoLevel",
			config: logger.LogConfig{
				Level:  "info",
				Format: "json",
			},
			expectedLevel:  zerolog.InfoLevel,
			expectedFormat: "json",
		},
		{
			name: "ConsoleFormatDebugLevel",
			config: logger.LogConfig{
				Level:  "debug",
				Format: "console",
			},
			expectedLevel:  zerolog.DebugLevel,
			expectedFormat: "console",
		},
		{
			name: "InvalidLevelDefaultsToInfo",
			config: logger.LogConfig{
				Level:  "invalid",
				Format: "json",
			},
			expectedLevel:  zerolog.InfoLevel,
			expectedFormat: "json",
		},
		{
			name: "EmptyLevelDefaultsToInfo",
			config: logger.LogConfig{
				Level:  "",
				Format: "json",
			},
			expectedLevel:  zerolog.InfoLevel,
			expectedFormat: "json",
		},
		{
			name: "WarnLevel",
			config: logger.LogConfig{
				Level:  "warn",
				Format: "json",
			},
			expectedLevel:  zerolog.WarnLevel,
			expectedFormat: "json",
		},
		{
			name: "ErrorLevel",
			config: logger.LogConfig{
				Level:  "error",
				Format: "json",
			},
			expectedLevel:  zerolog.ErrorLevel,
			expectedFormat: "json",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			originalLogger := log.Logger
			defer func() {
				log.Logger = originalLogger
			}()

			logger.Init(tc.config)

			require.Equal(t, tc.expectedLevel, log.Logger.GetLevel())
		})
	}
}

func TestLoggingFunctions(t *testing.T) {
	var buf bytes.Buffer
	originalLogger := log.Logger
	defer func() {
		log.Logger = originalLogger
	}()

	log.Logger = zerolog.New(&buf).Level(zerolog.DebugLevel)

	testCases := []struct {
		name     string
		logFunc  func()
		expected string
	}{
		{
			name: "Info",
			logFunc: func() {
				logger.Info("test info message")
			},
			expected: "test info message",
		},
		{
			name: "Debug",
			logFunc: func() {
				logger.Debug("test debug message")
			},
			expected: "test debug message",
		},
		{
			name: "Warn",
			logFunc: func() {
				logger.Warn("test warn message")
			},
			expected: "test warn message",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buf.Reset()
			tc.logFunc()
			require.Contains(t, buf.String(), tc.expected)
		})
	}
}

func TestFormattedLoggingFunctions(t *testing.T) {
	var buf bytes.Buffer
	originalLogger := log.Logger
	defer func() {
		log.Logger = originalLogger
	}()

	log.Logger = zerolog.New(&buf).Level(zerolog.DebugLevel)

	testCases := []struct {
		name     string
		logFunc  func()
		expected string
	}{
		{
			name: "Infof",
			logFunc: func() {
				logger.Infof("test info %s %d", "message", 123)
			},
			expected: "test info message 123",
		},
		{
			name: "Debugf",
			logFunc: func() {
				logger.Debugf("test debug %s %d", "message", 456)
			},
			expected: "test debug message 456",
		},
		{
			name: "Warnf",
			logFunc: func() {
				logger.Warnf("test warn %s %d", "message", 789)
			},
			expected: "test warn message 789",
		},
		{
			name: "Errorf",
			logFunc: func() {
				logger.Errorf("test error %s %d", "message", 999)
			},
			expected: "test error message 999",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buf.Reset()
			tc.logFunc()
			require.Contains(t, buf.String(), tc.expected)
		})
	}
}

func TestErrorLogging(t *testing.T) {
	var buf bytes.Buffer
	originalLogger := log.Logger
	defer func() {
		log.Logger = originalLogger
	}()

	log.Logger = zerolog.New(&buf).Level(zerolog.ErrorLevel)

	testCases := []struct {
		name     string
		logFunc  func()
		expected string
	}{
		{
			name: "ErrorWithString",
			logFunc: func() {
				logger.Error("string error message")
			},
			expected: "string error message",
		},
		{
			name: "ErrorWithError",
			logFunc: func() {
				logger.Error(os.ErrNotExist)
			},
			expected: "file does not exist",
		},
		{
			name: "ErrorWithOther",
			logFunc: func() {
				logger.Error(123)
			},
			expected: "123",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buf.Reset()
			tc.logFunc()
			require.Contains(t, buf.String(), tc.expected)
		})
	}
}

func TestLogConfigStruct(t *testing.T) {
	config := logger.LogConfig{
		Level:  "debug",
		Format: "console",
	}

	require.Equal(t, "debug", config.Level)
	require.Equal(t, "console", config.Format)
}
