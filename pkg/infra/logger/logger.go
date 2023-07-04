package logger

import (
	"fmt"
	"log"
	"os"
)

// ILogger ILogger
type ILogger interface {
	Log(msg string) error
	Error(msg string) error
	New() (ILogger, error)
}

// Logger Logger
type Logger struct {
	RootPath string
}

// Log Log
func (l *Logger) Log(msg string) error {
	fmt.Println(msg)
	return logFile(msg, l.RootPath+l.InfoBasePath())
}

// Error Error
func (l *Logger) Error(msg string) error {
	fmt.Println(msg)
	return errorFile(msg, l.RootPath+l.ErrorBasePath())
}

// New New
func (l Logger) New() (ILogger, error) {
	return &Logger{RootPath: l.RootPath}, nil
}

// InfoBasePath InfoBasePath
func (l *Logger) InfoBasePath() string {
	return "var/log/info.log"
}

// ErrorBasePath ErrorBasePath
func (l *Logger) ErrorBasePath() string {
	return "var/log/error.log"
}

func logFile(msg string, fileLog string) error {
	file, err := os.OpenFile(fileLog, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	log.SetOutput(file)
	log.Println(msg)
	return nil
}

func errorFile(msg string, fileLog string) error {
	file, err := os.OpenFile(fileLog, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	log.SetOutput(file)
	log.Println(msg)
	return nil
}
