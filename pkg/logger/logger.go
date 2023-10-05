package logger

import (
	"go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {
	// TODO: add environment variables into the raw.json config, and also load them from a .yaml file.
	// TODO: add file logging to root folder logs/

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.TimeKey = "timestamp |"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	config.OutputPaths = []string{
		"stderr",
	}
	config.ErrorOutputPaths = []string{
		"stderr",
	}


	return zap.Must(config.Build())
}
