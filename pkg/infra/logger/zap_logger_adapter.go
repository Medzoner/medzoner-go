package logger

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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
func NewLoggerAdapter(useSugar UseSugar) *ZapLoggerAdapter {
	zl := ZapLoggerAdapter{
		UseSugar: bool(useSugar),
	}
	logger := zl.New()
	return logger
}

// New New
func (z ZapLoggerAdapter) New() *ZapLoggerAdapter {
	rawJSON := []byte(`{
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

	cfg := zap.Config{}
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	cfg.EncoderConfig.EncodeLevel = func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + level.CapitalString() + "]")
	}
	zapLogger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	defer z.deferLogger(zapLogger)

	z.Zap = zapLogger
	return &z
}

func (z ZapLoggerAdapter) deferLogger(zapLogger *zap.Logger) {
	err := zapLogger.Sync()
	if err == nil {
		fmt.Println("sync defer")
	}
}

// Log Log
func (z ZapLoggerAdapter) Log(msg string) {
	if z.UseSugar {
		sugar := z.Zap.Sugar()
		sugar.Infow(msg)
	}
	if !z.UseSugar {
		z.Zap.Info(msg)
	}
}

// Error Error
func (z ZapLoggerAdapter) Error(msg string) {
	if z.UseSugar {
		sugar := z.Zap.Sugar()
		sugar.Errorw(msg)
	}
	if !z.UseSugar {
		z.Zap.Error(msg)
	}
}
