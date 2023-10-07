package logger

import (
	"os"
	"time"

	"github.com/r00tk3y/prying-deep/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {
	config := configs.GetConfig()

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     customEncoderTime,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// default encoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	var level zapcore.Level
	switch config.LoggerConf.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	if config.LoggerConf.Encoder == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}
	//TODO: fix writing to file later
	//
	//fileWriter := zapcore.AddSync(&lumberjack.Logger{
	//	Filename:   fmt.Sprintf("provider-%s.log", time.Now().Format("2006-01-02")),
	//	MaxSize:    5, // megabytes
	//	MaxBackups: 3,
	//	MaxAge:     28, // days
	//})
	//
	//fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stderr), level),
		//zapcore.NewCore(fileEncoder, fileWriter, level),
	)

	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func customEncoderTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Infof(template string, args ...any) {
	Logger.Sugar().Infof(template, args...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Warnf(template string, args ...any) {
	Logger.Sugar().Warnf(template, args...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Errorf(template string, args ...any) {
	Logger.Sugar().Errorf(template, args...)
}

func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func Debugf(template string, args ...any) {
	Logger.Sugar().Debugf(template, args...)
}

func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}

func Fatalf(template string, args ...any) {
	Logger.Sugar().Fatalf(template, args...)
}

func Panic(msg string, fields ...zap.Field) {
	Logger.Panic(msg, fields...)
}

func Panicf(template string, args ...any) {
	Logger.Sugar().Panicf(template, args...)
}
