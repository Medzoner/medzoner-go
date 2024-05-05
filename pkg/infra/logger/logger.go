package logger

// ILogger ILogger
type ILogger interface {
	Log(msg string)
	Error(msg string)
}
