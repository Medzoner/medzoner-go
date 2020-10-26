package logger_test

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"gotest.tools/assert"
	"io/ioutil"
	"regexp"
	"testing"
)

func TestLoggerSuccess(t *testing.T) {
	t.Run("Unit: test logger log success", func(t *testing.T) {
		loggerTest := &logger.Logger{RootPath: "../../../"}
		loggerInstanceTest := loggerTest.New()
		message := "good test"
		loggerInstanceTest.Log(message)

		b, err := ioutil.ReadFile("../../../" + loggerTest.InfoBasePath())
		if err != nil {
			panic(err)
		}
		isExist, err := regexp.Match(message, b)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, isExist, true)
	})
}

func TestLoggerErrorSuccess(t *testing.T) {
	t.Run("Unit: test logger log error success", func(t *testing.T) {
		loggerTest := logger.Logger{RootPath: "../../../"}
		message := "this is a log error"
		loggerTest.Error(message)

		b, err := ioutil.ReadFile("../../../" + loggerTest.ErrorBasePath())
		if err != nil {
			panic(err)
		}
		isExist, err := regexp.Match(message, b)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, isExist, true)
	})
}

//func TestLoggerFailed(t *testing.T) {
//	t.Run("Unit: test logger log failed with bad RootPath", func(t *testing.T) {
//		loggerTest := logger.Logger{RootPath: "../../"}
//		loggerInstanceTest := loggerTest.New()
//		message := "this is a log error"
//
//		defer func() {
//			if r := recover(); r == nil {
//				t.Errorf("The code did not panic")
//			}
//		}()
//		loggerInstanceTest.Error(message)
//	})
//}
