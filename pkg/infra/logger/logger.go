package logger

import (
	"fmt"
	"log"
	"os"
)

type ILogger interface {
	Log(msg string)
	Error(msg string)
	New() ILogger
}

type Logger struct {
	RootPath string
}

func (l *Logger) Log(msg string) {
	n, e := fmt.Println(msg)
	logFile(msg, l.RootPath+l.InfoBasePath())
	if e != nil {
		os.Exit(1)
	}
	if n > 0 {
		return
	}
}

func (l *Logger) Error(msg string) {
	fmt.Println(msg)
	errorFile(msg, l.RootPath+l.ErrorBasePath())
}

func (l Logger) New() ILogger {
	return &Logger{}
}

func (l *Logger) InfoBasePath() string {
	return "var/log/info.log"
}

func (l *Logger) ErrorBasePath() string {
	return "var/log/error.log"
}

func logFile(msg string, fileLog string) {
	file, err := os.OpenFile(fileLog, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = file.Close()
	}()
	log.SetOutput(file)
	log.Println(msg)
}

func errorFile(msg string, fileLog string) {
	file, err := os.OpenFile(fileLog, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = file.Close()
	}()
	log.SetOutput(file)
	log.Println(msg)
}
