package logger_test

import (
	"testing"

	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
)

func TestZapLoggerAdapterSuccess(t *testing.T) {
	t.Run("Unit: test ZapLoggerAdapter log success", func(t *testing.T) {
		loggerTest, err := logger.NewLoggerAdapter(config.Config{DebugMode: true})
		if err != nil {
			t.Errorf("error creating logger: %v", err)
		}
		loggerTest.Log("log zap")
	})
	t.Run("Unit: test ZapLoggerAdapter error log success", func(t *testing.T) {
		loggerTest, err := logger.NewLoggerAdapter(config.Config{DebugMode: true})
		if err != nil {
			t.Errorf("error creating logger: %v", err)
		}
		loggerTest.Error("error zap")
	})
}
