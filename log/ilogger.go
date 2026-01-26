package log

type ILogger interface {
	Print(log *Log)
	Info(msg string)
	Debug(msg string)
	Error(msg string)
	Warning(msg string)
}
