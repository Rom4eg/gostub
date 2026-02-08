package template

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateContext_Base64Encode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "PASS: base64 encode",
			str:  "foo",
			want: "Zm9v",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()
			ctx := NewTemplateContext(req, res)
			got, err := ctx.Base64Encode(tt.str)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
