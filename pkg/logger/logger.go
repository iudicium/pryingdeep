package logger

import (
	"go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)


var Logger *zap.Logger


func InitLogger()  {
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


	Logger = zap.Must(config.Build())
}

// Maybe i omly need sugared llogging for now
func Debug(message string, fields ...interface{}) {
    Logger.Sugar().Debugw(message, fields)
}

func Info(message string, fields ...interface{}) {
    Logger.Sugar().Infow(message, fields...)
}

func Warn(message string, fields ...interface{}) {
    Logger.Sugar().Warnw(message, fields...)
}

func Error(message string, fields ...interface{}) {
    Logger.Sugar().Errorw(message, fields...)
}

func Fatal(message string, fields ...interface{}) {
    Logger.Sugar().Fatalw(message, fields...)
}
