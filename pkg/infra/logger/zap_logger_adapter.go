package logger

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapConfig = []byte(`{
		"level": "debug",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoding": "json",
		"encoderConfig": {
			"timeKey": "ts",
			"timeEncoder": "ISO8601",
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		}
	}`)

type UseSugar bool

func NewUseSugar() UseSugar {
	return false
}

// ZapLoggerAdapter ZapLoggerAdapter
type ZapLoggerAdapter struct {
	Zap      *zap.Logger
	UseSugar bool
}

// NewLoggerAdapter NewLoggerAdapter
func NewLoggerAdapter(useSugar UseSugar) (*ZapLoggerAdapter, error) {
	cfg := zap.Config{}
	if err := json.Unmarshal(zapConfig, &cfg); err != nil {
		return nil, err
	}

	cfg.EncoderConfig.EncodeLevel = func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + level.CapitalString() + "]")
	}
	zapLogger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	zl := ZapLoggerAdapter{
		UseSugar: bool(useSugar),
		Zap:      zapLogger,
	}
	defer zl.deferLogger(zapLogger)

	return &zl, nil
}

func (z ZapLoggerAdapter) deferLogger(zapLogger *zap.Logger) {
	if err := zapLogger.Sync(); err == nil {
		fmt.Println("sync defer")
	}
}

// Log Log
func (z ZapLoggerAdapter) Log(msg string) {
	if z.UseSugar {
		sugar := z.Zap.Sugar()
		sugar.Infow(msg)
	}
	z.Zap.Info(msg)
}

// Error Error
func (z ZapLoggerAdapter) Error(msg string) {
	if z.UseSugar {
		z.Zap.Sugar().Errorw(msg)
		return
	}
	z.Zap.Error(msg)
}
