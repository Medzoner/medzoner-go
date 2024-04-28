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
	logger, err := zl.New()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return logger
}

// New New
func (z ZapLoggerAdapter) New() (*ZapLoggerAdapter, error) {
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
		return nil, err
	}

	cfg.EncoderConfig.EncodeLevel = func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + level.CapitalString() + "]")
	}
	zapLogger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	defer z.deferLogger(zapLogger)
	z.Zap = zapLogger
	return &z, nil
}

func (z ZapLoggerAdapter) deferLogger(zapLogger *zap.Logger) {
	err := zapLogger.Sync()
	if err == nil {
		fmt.Println("sync defer")
	}
}

// Log Log
func (z ZapLoggerAdapter) Log(msg string) error {
	if z.UseSugar {
		sugar := z.Zap.Sugar()
		sugar.Infow(msg)
	}
	if !z.UseSugar {
		z.Zap.Info(msg)
	}
	return nil
}

// Error Error
func (z ZapLoggerAdapter) Error(msg string) error {
	if z.UseSugar {
		sugar := z.Zap.Sugar()
		sugar.Errorw(msg)
	}
	if !z.UseSugar {
		z.Zap.Error(msg)
	}
	return nil
}
