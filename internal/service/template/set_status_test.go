package template

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateContext_SetStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
	}{
		{
			name:   "PASS: set status",
			status: http.StatusNotImplemented,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()
			ctx := NewTemplateContext(req, res)
			s, err := ctx.SetStatus(tt.status)
			assert.NoError(t, err)
			assert.Equal(t, "", s)
			assert.Equal(t, tt.status, ctx.status)
		})
	}
}
