package logs

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Std = NewLogging("dynamic")

func NewLogging(service string) *zap.SugaredLogger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var atom zap.AtomicLevel
	Env := os.Getenv("Env")
	if Env == "prod" {
		atom = zap.NewAtomicLevelAt(zap.WarnLevel)
	} else {
		atom = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	config := zap.Config{
		Level:             atom,
		DisableCaller:     false,
		Development:       false,
		DisableStacktrace: false,
		Encoding:          "console",
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
	}
	config.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	var err error
	log, err := config.Build()
	log = log.Named(service)
	if err != nil {
		panic(fmt.Sprintf("log 初始化失败: %v", err))
	}
	return log.Sugar()
}
