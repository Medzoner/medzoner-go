package logger_test

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"testing"
)

func TestZapLoggerAdapterSuccess(t *testing.T) {
	t.Run("Unit: test ZapLoggerAdapter log success", func(t *testing.T) {
		loggerTest := logger.ZapLoggerAdapter{
			UseSugar: false,
		}.New()
		loggerTest.Log("log zap")
	})
	t.Run("Unit: test ZapLoggerAdapter error log success", func(t *testing.T) {
		loggerTest := logger.NewLoggerAdapter(false).New()
		loggerTest.Error("error zap")
	})
	t.Run("Unit: test ZapLoggerAdapter (sugared) log success", func(t *testing.T) {
		loggerTest := logger.NewLoggerAdapter(true).New()
		loggerTest.Log("log (sugared) zap")
	})
	t.Run("Unit: test ZapLoggerAdapter (sugared) error log success", func(t *testing.T) {
		loggerTest := logger.NewLoggerAdapter(true).New()
		loggerTest.Error("error (sugared) zap")
	})
}
