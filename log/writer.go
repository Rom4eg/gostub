package log

import (
	"io"
	"os"
)

var writer io.Writer = os.Stdout

func SetWriter(w io.Writer) {
	writer = w
}
