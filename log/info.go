package log

func (l *Logger) Info(msg string) {
	if level < LevelInfo {
		return
	}

	_l := New(l.prefix, msg, LevelInfo)
	l.Print(_l)
}

func Info(msg string) {
	l := Logger{}
	l.Info(msg)
}
