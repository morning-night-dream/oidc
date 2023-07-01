package log

import (
	"os"

	"go.uber.org/zap"
)

func Log() *zap.Logger {
	opts := []zap.Option{}

	config := zap.NewProductionConfig()
	if os.Getenv("DEBUG_MODE") != "" {
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	logger, _ := config.Build(opts...)

	return logger
}
