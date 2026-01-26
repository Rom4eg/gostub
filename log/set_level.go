package log

func SetLevel(l LogLevel) {
	level = l
}

func SetLevelS(s string) {
	switch s {
	case "info":
		SetLevel(LevelInfo)
	case "error":
		SetLevel(LevelError)
	case "warning":
		SetLevel(LevelWarning)
	case "debug":
		SetLevel(LevelDebug)
	case "none":
		SetLevel(LevelNone)
	default:
		SetLevel(LevelUndefined)
	}
}
