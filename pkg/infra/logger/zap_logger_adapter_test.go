package logger_test

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"testing"
)

func TestZapLoggerAdapterSuccess(t *testing.T) {
	t.Run("Unit: test ZapLoggerAdapter log success", func(t *testing.T) {
		loggerTest, err := logger.NewLoggerAdapter(false)
		if err != nil {
			t.Errorf("error creating logger: %v", err)
		}
		loggerTest.Log("log zap")
	})
	t.Run("Unit: test ZapLoggerAdapter error log success", func(t *testing.T) {
		loggerTest, err := logger.NewLoggerAdapter(false)
		if err != nil {
			t.Errorf("error creating logger: %v", err)
		}
		loggerTest.Error("error zap")
	})
	t.Run("Unit: test ZapLoggerAdapter (sugared) log success", func(t *testing.T) {
		loggerTest, err := logger.NewLoggerAdapter(true)
		if err != nil {
			t.Errorf("error creating logger: %v", err)
		}
		loggerTest.Log("log (sugared) zap")
	})
	t.Run("Unit: test ZapLoggerAdapter (sugared) error log success", func(t *testing.T) {
		loggerTest, err := logger.NewLoggerAdapter(true)
		if err != nil {
			t.Errorf("error creating logger: %v", err)
		}
		loggerTest.Error("error (sugared) zap")
	})
}
