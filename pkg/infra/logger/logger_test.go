package logger_test

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"gotest.tools/assert"
	"io/ioutil"
	"regexp"
	"testing"
)

func TestLoggerSuccess(t *testing.T) {
	t.Run("Unit: test logger log success", func(t *testing.T) {
		loggerTest := &logger.Logger{RootPath: "../../../"}
		loggerInstanceTest, _ := loggerTest.New()
		message := "good test"
		err := loggerInstanceTest.Log(message)
		if err != nil {
			fmt.Println(err)
		}

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
		_ = loggerTest.Error(message)

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

func TestLoggerFailed(t *testing.T) {
	t.Run("Unit: test logger log error failed with bad RootPath", func(t *testing.T) {
		loggerTest := logger.Logger{RootPath: "../../"}
		loggerInstanceTest, _ := loggerTest.New()
		message := "this is a log error failed"

		_ = loggerInstanceTest.Error(message)
	})
	t.Run("Unit: test logger log failed with bad RootPath", func(t *testing.T) {
		loggerTest := logger.Logger{RootPath: "../../"}
		loggerInstanceTest, _ := loggerTest.New()
		message := "this is a log failed"

		err := loggerInstanceTest.Log(message)
		if err != nil {
			fmt.Println(err)
		}
	})
}
