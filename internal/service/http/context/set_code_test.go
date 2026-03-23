package context

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStubContext_SetCode(t *testing.T) {
	t.Parallel()

	ctx := &StubContext{}
	v, err := ctx.SetCode(200)
	assert.NoError(t, err)
	assert.Equal(t, "", v)
	assert.Equal(t, ctx.code, 200)
}
