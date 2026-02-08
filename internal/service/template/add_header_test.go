package template

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateContext_AddHeader(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		headers map[string]string
	}{
		{
			name: "PASS: add header",
			headers: map[string]string{
				"foo": "bar",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()
			ctx := NewTemplateContext(req, res)
			for k, v := range tt.headers {
				s, e := ctx.AddHeader(k, v)
				assert.NoError(t, e)
				assert.Equal(t, "", s)

				actual := res.Header().Get(k)
				assert.Equal(t, v, actual)
			}
		})
	}
}
