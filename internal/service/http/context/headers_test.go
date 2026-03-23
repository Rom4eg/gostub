package context

import (
	"net/http"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestStubContext_Headers(t *testing.T) {
	t.Parallel()

	ctx := &StubContext{
		headers: http.Header{
			"foo": []string{"bar"},
		},
	}
	v := ctx.Headers()
	assert.Equal(t, v, ctx.headers)
}
