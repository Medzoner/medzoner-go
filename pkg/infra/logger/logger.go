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
	logFile(msg, l.RootPath + "var/log/info.log")
	if e != nil {
		os.Exit(1)
	}
	if n > 0 {
		return
	}
}

func (l *Logger) Error(msg string) {
	fmt.Println(msg)
	errorFile(msg, l.RootPath + "var/log/error.log")
}

func (l Logger) New() ILogger {
	return &Logger{}
}

func logFile(msg string, fileLog string) {
	file, err := os.OpenFile(fileLog, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println(msg)
}

func errorFile(msg string, fileLog string)  {
	file, err := os.OpenFile(fileLog, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println(msg)
}