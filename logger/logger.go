package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var (
	Logger *zap.Logger
	once   sync.Once
)

// InitLogger initializes the global zap logger
func Init() {
	once.Do(func() {
		config := zap.Config{
			Encoding:         "json",
			Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "time",
				LevelKey:       "level",
				MessageKey:     "msg",
				CallerKey:      "caller",
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeLevel:    zapcore.CapitalLevelEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
			},
		}

		var err error
		Logger, err = config.Build()
		if err != nil {
			panic(err)
		}
	})
}

// Sync flushes any buffered logger entries
func Sync() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}
