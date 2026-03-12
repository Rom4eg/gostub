package funcs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSl(t *testing.T) {

	t.Run("PASS: sleep", func(t *testing.T) {
		start := time.Now()
		howLong := 5
		ms, e := Sleep(howLong)
		assert.NoError(t, e)
		assert.Equal(t, howLong, ms)

		end := time.Now()
		if end.Sub(start) < 5*time.Millisecond {
			t.Errorf("expected %dms, got %v", howLong, end.Sub(start))
		}
	})
}
