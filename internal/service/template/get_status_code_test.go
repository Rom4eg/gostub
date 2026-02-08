package template

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateContext_GetStatusCode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		want   int
		status int
	}{
		{
			name:   "PASS: get status",
			want:   http.StatusNotImplemented,
			status: http.StatusNotImplemented,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()
			ctx := NewTemplateContext(req, res)
			ctx.status = tt.status

			got := ctx.GetStatusCode()
			assert.Equal(t, tt.want, got)
		})
	}
}
