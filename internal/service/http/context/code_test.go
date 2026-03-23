package context

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestStubContext_Code(t *testing.T) {
	t.Parallel()

	ctx := &StubContext{
		code: 200,
	}

	c := ctx.Code()
	assert.Equal(t, 200, c)
}
