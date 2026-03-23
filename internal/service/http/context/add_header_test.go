package context

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStubContext_AddHeader(t *testing.T) {
	t.Parallel()

	ctx := &StubContext{
		headers: make(http.Header),
	}
	v, err := ctx.AddHeader("foo", "bar")
	assert.NoError(t, err)
	assert.Equal(t, "", v)
	assert.Equal(t, http.Header{"Foo": []string{"bar"}}, ctx.headers)
}
