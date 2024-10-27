package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLoggerAdapter ZapLoggerAdapter
type ZapLoggerAdapter struct {
	Zap *zap.Logger
}

// NewLoggerAdapter NewLoggerAdapter
func NewLoggerAdapter() (*ZapLoggerAdapter, error) {
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		Encoding:         "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:    "ts",
			EncodeTime: zapcore.ISO8601TimeEncoder,
			MessageKey: "message",
			LevelKey:   "level",
			EncodeLevel: func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString("[" + level.CapitalString() + "]")
			},
		},
	}

	zapLogger, err := cfg.Build()

	zl := ZapLoggerAdapter{
		Zap: zapLogger,
	}
	defer zapLogger.Sync()

	return &zl, err
}

// Log Log
func (z ZapLoggerAdapter) Log(msg string) {
	z.Zap.Info(msg)
}

// Error Error
func (z ZapLoggerAdapter) Error(msg string) {
	z.Zap.Error(msg)
}
