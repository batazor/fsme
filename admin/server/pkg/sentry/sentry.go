package sentry

import (
	"bytes"
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
)

func init() {
	// Logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Set config from .ENV
	if _, err := os.Stat(".env"); err == nil {
		config, err := ioutil.ReadFile(".env")
		if err != nil {
			logger.Error("Error read .env file", zap.Error(err))
		}

		viper.SetConfigType("env")
		viper.ReadConfig(bytes.NewBuffer(config))
	}
}

func Run() {
	// Logger ==================================================================
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Load Sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn: viper.GetString("SENTRY_DSN"),
	})

	if err != nil {
		logger.Error("Sentry initialization failed", zap.Error(err))
	}
}
