package context

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStubContext_DeleteHeader(t *testing.T) {
	t.Parallel()

	ctx := &StubContext{
		headers: http.Header{
			"Foo": []string{"bar"},
		},
	}
	v, err := ctx.DeleteHeader("foo")
	assert.NoError(t, err)
	assert.Equal(t, "", v)
	assert.Equal(t, http.Header{}, ctx.headers)
}
