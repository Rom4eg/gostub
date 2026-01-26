package log

import (
	"fmt"
	"io"
	"strings"
)

func (l *Logger) Print(log *Log) {
	msg := fmt.Sprintf("%s\n", log.String())
	b := strings.NewReader(msg)
	_, err := io.Copy(writer, b)
	if err != nil {
		panic(err)
	}
}
