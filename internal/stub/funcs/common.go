package funcs

import "time"

func Sleep(ms int) (int, error) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
	return ms, nil
}
