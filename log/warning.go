package log

func (l *Logger) Warning(msg string) {
	if level < LevelWarning {
		return
	}

	_l := New(l.prefix, msg, LevelWarning)
	l.Print(_l)
}

func Warning(msg string) {
	l := Logger{}
	l.Warning(msg)
}
