package flags

import "os"

func Get() *Flags {
	if f == nil {
		Parse(os.Args[1:])
	}
	return f
}
