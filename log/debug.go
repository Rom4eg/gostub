package log

func (l *Logger) Debug(msg string) {
	if level < LevelDebug {
		return
	}

	_l := New(l.prefix, msg, LevelDebug)
	l.Print(_l)
}

func Debug(msg string) {
	l := Logger{}
	l.Debug(msg)
}
