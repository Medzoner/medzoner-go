package logger

// ILogger ILogger
type ILogger interface {
	Log(msg string) error
	Error(msg string) error
}
