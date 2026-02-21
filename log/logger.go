package log

//go:generate mockgen -typed -destination=mocks/logger.go -package=mocks . ILogger

type ILogger interface {
	Print(log *Log)
	Info(msg string)
	Debug(msg string)
	Error(msg string)
	Warning(msg string)
}

type Logger struct {
	prefix string
}

func NewLogger(p string) *Logger {
	return &Logger{
		prefix: p,
	}
}
