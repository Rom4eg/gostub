package template

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateContext_GetStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		want   string
		status int
	}{
		{
			name:   "PASS: get status",
			want:   fmt.Sprint(http.StatusNotImplemented),
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

			got, e := ctx.GetStatus()
			assert.NoError(t, e)
			assert.Equal(t, tt.want, got)
		})
	}
}
