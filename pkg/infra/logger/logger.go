package logger

type ILogger interface {
	Log(msg string) error
	Error(msg error) error
}
