package log

import "fmt"

func (l *Logger) With(msg string) *Logger {
	p := msg
	if l.prefix != "" {
		p = fmt.Sprintf("%s %s", l.prefix, msg)
	}
	return &Logger{
		prefix: p,
	}
}

func With(msg string) *Logger {
	return &Logger{
		prefix: msg,
	}
}
