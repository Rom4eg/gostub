package context

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStubContext_JsonBody(t *testing.T) {
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
				buf.WriteString(`{"foo": "bar"}`)
				return io.NopCloser(&buf)
			},
			expect: func(t *testing.T, b map[string]any, err error) {
				assert.Equal(t, map[string]any{"foo": "bar"}, b)
				assert.NoError(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := http.Request{
				Body: tt.getBody(),
			}

			ctx := &StubContext{
				Request: &req,
			}
			b, err := ctx.JsonBody()
			tt.expect(t, b, err)
		})
	}
}
