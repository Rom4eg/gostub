package context

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStubContext_MultipartForm(t *testing.T) {
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
			name: "FAIL: invalid form",
			getBody: func() io.ReadCloser {
				var buf bytes.Buffer
				buf.WriteString("invalid;form")
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
				buf.WriteString("--------------------------vyljaNc3fiOTtzT3jZmoCJ\r\nContent-Disposition: form-data; name=\"foo\"\r\n\r\nbar\r\n--------------------------vyljaNc3fiOTtzT3jZmoCJ--\r\n")
				return io.NopCloser(&buf)
			},
			expect: func(t *testing.T, b map[string]any, err error) {
				assert.Equal(t, map[string]any{"foo": []string{"bar"}}, b)
				assert.NoError(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := http.Request{
				Body:   tt.getBody(),
				Method: http.MethodPost,
				Header: http.Header{
					"Content-Type": []string{"multipart/form-data; boundary=------------------------vyljaNc3fiOTtzT3jZmoCJ"},
				},
			}

			ctx := &StubContext{
				Request: &req,
			}

			b, err := ctx.MultipartBody()
			tt.expect(t, b, err)
		})
	}
}
