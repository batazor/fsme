package sentry

import (
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

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
