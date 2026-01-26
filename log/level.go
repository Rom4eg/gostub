package log

import "encoding/json"

type LogLevel int

const (
	LevelUndefined LogLevel = -1
	LevelNone      LogLevel = iota
	LevelInfo
	LevelError
	LevelWarning
	LevelDebug
)

var level LogLevel = LevelInfo

func (l LogLevel) String() string {
	switch l {
	case LevelInfo:
		return "info"
	case LevelError:
		return "error"
	case LevelWarning:
		return "warning"
	case LevelDebug:
		return "debug"
	default:
		return "undefined"
	}
}

func (l LogLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}
