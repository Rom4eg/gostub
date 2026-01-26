package log

import "fmt"

func (l *Logger) With(msg string) *Logger {
	return &Logger{
		prefix: fmt.Sprintf("%s %s", l.prefix, msg),
	}
}

func With(msg string) *Logger {
	return &Logger{
		prefix: msg,
	}
}
