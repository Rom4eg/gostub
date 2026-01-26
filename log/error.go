package log

func (l *Logger) Error(msg string) {
	if level < LevelError {
		return
	}

	_l := New(l.prefix, msg, LevelError)
	l.Print(_l)
}

func Error(msg string) {
	l := Logger{}
	l.Error(msg)
}
