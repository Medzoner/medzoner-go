package logger_test

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"testing"
)

func TestZapLoggerAdapterSuccess(t *testing.T) {
	t.Run("Unit: test ZapLoggerAdapter log success", func(t *testing.T) {
		loggerTest, _ := logger.ZapLoggerAdapter{
			UseSugar: false,
		}.New()
		loggerTest.Log("log zap")
	})
	t.Run("Unit: test ZapLoggerAdapter error log success", func(t *testing.T) {
		loggerTest, _ := logger.ZapLoggerAdapter{
			UseSugar: false,
		}.New()
		loggerTest.Error("error zap")
	})
	t.Run("Unit: test ZapLoggerAdapter (sugared) log success", func(t *testing.T) {
		loggerTest, _ := logger.ZapLoggerAdapter{
			UseSugar: true,
		}.New()
		loggerTest.Log("log (sugared) zap")
	})
	t.Run("Unit: test ZapLoggerAdapter (sugared) error log success", func(t *testing.T) {
		loggerTest, err := logger.ZapLoggerAdapter{
			UseSugar: true,
		}.New()
		if err != nil {
			fmt.Println(err)
		}
		loggerTest.Error("error (sugared) zap")
	})
}
