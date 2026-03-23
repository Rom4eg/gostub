package context

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStubContext_Body(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		getBody func() io.ReadCloser
		expect  func(*testing.T, map[string]any, error)
	}{
		{
			name: "FAIL: empty body",
			getBody: func() io.ReadCloser {
				return nil
			},
			expect: func(t *testing.T, b map[string]any, err error) {
				assert.Equal(t, map[string]any{}, b)
				assert.NoError(t, err)
			},
		},

		{
			name: "FAIL: invalid json",
			getBody: func() io.ReadCloser {
				var buf bytes.Buffer
				buf.WriteString("invalid json")
				return io.NopCloser(&buf)
			},
			expect: func(t *testing.T, b map[string]any, err error) {
				assert.Equal(t, map[string]any{}, b)
				assert.Error(t, err)
			},
		},

		{
			name: "PASS: OK",
			getBody: func() io.ReadCloser {
				var buf bytes.Buffer
				buf.WriteString(`{"test": "test"}`)
				return io.NopCloser(&buf)
			},
			expect: func(t *testing.T, b map[string]any, err error) {
				assert.Equal(t, map[string]any{"test": "test"}, b)
				assert.NoError(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := http.Request{
				Body: tt.getBody(),
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
				Method: http.MethodPost,
			}

			ctx := New(&req)
			b, err := ctx.Body()
			tt.expect(t, b, err)
		})
	}
}
