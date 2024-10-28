package logger

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLoggerAdapter ZapLoggerAdapter
type ZapLoggerAdapter struct {
	Zap    *zap.Logger
	Config config.Config
}

// NewLoggerAdapter NewLoggerAdapter
func NewLoggerAdapter(config config.Config) (*ZapLoggerAdapter, error) {
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.ErrorLevel),
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

	if config.DebugMode {
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.Level.SetLevel(zap.DebugLevel)
	}

	zapLogger, err := cfg.Build()
	defer zapLogger.Sync()

	return &ZapLoggerAdapter{
		Zap: zapLogger,
	}, err
}

// Log Log
func (z ZapLoggerAdapter) Log(msg string) {
	z.Zap.Info(msg)
}

// Error Error
func (z ZapLoggerAdapter) Error(msg string) {
	z.Zap.Error(msg)
}
