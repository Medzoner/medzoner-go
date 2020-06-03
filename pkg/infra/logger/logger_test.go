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
		message := "good test"
		loggerTest.Log(message)

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

