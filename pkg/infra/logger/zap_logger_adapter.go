package logger

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLoggerAdapter struct {
	RootPath string
	Zap      *zap.Logger
	UseSugar bool
}

func (z ZapLoggerAdapter) New() ILogger {
	rawJSON := []byte(`{
		"level": "debug",
		"outputPaths": ["stdout", "` + z.RootPath + `var/log/info.log"],
		"errorOutputPaths": ["stderr", "` + z.RootPath + `var/log/error.log"],
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

func (z ZapLoggerAdapter) Log(msg string) {
	if z.UseSugar {
		sugar := z.Zap.Sugar()
		sugar.Infow(msg)
	}
	if !z.UseSugar {
		z.Zap.Info(msg)
	}
}

func (z ZapLoggerAdapter) Error(msg string) {
	if z.UseSugar {
		sugar := z.Zap.Sugar()
		sugar.Errorw(msg)
	}
	if !z.UseSugar {
		z.Zap.Error(msg)
	}
}
