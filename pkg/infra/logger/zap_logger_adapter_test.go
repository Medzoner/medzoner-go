package logger_test

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"testing"
)

func TestZapLoggerAdapterSuccess(t *testing.T) {
	t.Run("Unit: test ZapLoggerAdapter log success", func(t *testing.T) {
		loggerTest, _ := logger.ZapLoggerAdapter{
			RootPath: "../../../",
			UseSugar: false,
		}.New()
		err := loggerTest.Log("log zap")
		if err != nil {
			fmt.Println(err)
		}
	})
	t.Run("Unit: test ZapLoggerAdapter error log success", func(t *testing.T) {
		loggerTest, _ := logger.ZapLoggerAdapter{
			RootPath: "../../../",
			UseSugar: false,
		}.New()
		_ = loggerTest.Error("error zap")
	})
	t.Run("Unit: test ZapLoggerAdapter (sugared) log success", func(t *testing.T) {
		loggerTest, _ := logger.ZapLoggerAdapter{
			RootPath: "../../../",
			UseSugar: true,
		}.New()
		err := loggerTest.Log("log (sugared) zap")
		if err != nil {
			fmt.Println(err)
		}
	})
	t.Run("Unit: test ZapLoggerAdapter (sugared) error log success", func(t *testing.T) {
		loggerTest, err := logger.ZapLoggerAdapter{
			RootPath: "../../../",
			UseSugar: true,
		}.New()
		if err != nil {
			fmt.Println(err)
		}
		_ = loggerTest.Error("error (sugared) zap")
	})
}
