package log

import (
	"encoding/json"
	"fmt"
	"time"
)

type Log struct {
	Time  string   `json:"time"`
	Level LogLevel `json:"level,string"`
	Msg   string   `json:"msg"`
}

func New(p, msg string, l LogLevel) *Log {
	return &Log{
		Time:  time.Now().Format(time.RFC3339),
		Level: l,
		Msg:   fmt.Sprintf("%s %s", p, msg),
	}
}

func (l Log) String() string {
	b, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(b)
}
