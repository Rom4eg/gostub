package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	lm "github.com/Rom4eg/gostub/log/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_StartResponse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		body    []byte
		status  int
		headers http.Header
		prepare func(*testing.T, *lm.MockILogger)
		expect  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:    "PASS: test auto status - not implemented",
			body:    make([]byte, 0),
			status:  0,
			headers: make(http.Header),
			prepare: func(t *testing.T, l *lm.MockILogger) {
				l.EXPECT().Debug(gomock.Any()).Times(2)
			},
			expect: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotImplemented, rr.Code)
			},
		},
		{
			name:    "PASS: test auto status - OK",
			body:    []byte("not empty"),
			status:  0,
			headers: make(http.Header),
			prepare: func(t *testing.T, l *lm.MockILogger) {
				l.EXPECT().Debug(gomock.Any()).Times(2)
			},
			expect: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, rr.Code)
			},
		},
		{
			name:    "PASS: test custom status",
			body:    []byte("not empty"),
			status:  http.StatusAccepted,
			headers: make(http.Header),
			prepare: func(t *testing.T, l *lm.MockILogger) {
				l.EXPECT().Debug(gomock.Any()).Times(2)
			},
			expect: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusAccepted, rr.Code)
			},
		},
		{
			name:   "PASS: test custom header",
			body:   []byte("not empty"),
			status: http.StatusOK,
			headers: http.Header{
				"foo": []string{"bar"},
			},
			prepare: func(t *testing.T, l *lm.MockILogger) {
				l.EXPECT().Debug(gomock.Any()).Times(2)
			},
			expect: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, rr.Code)
				assert.Equal(t, "bar", rr.Header().Get("foo"))
			},
		},
		{
			name:    "PASS: test body",
			body:    []byte("some content"),
			status:  http.StatusOK,
			headers: make(http.Header),
			prepare: func(t *testing.T, l *lm.MockILogger) {
				l.EXPECT().Debug(gomock.Any()).Times(2)
			},
			expect: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, rr.Code)
				assert.Equal(t, "some content", rr.Body.String())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			l := lm.NewMockILogger(ctrl)
			tt.prepare(t, l)
			rw := httptest.NewRecorder()
			s := &Service{
				l: l,
			}
			resp := &Response{
				Body:   tt.body,
				Code:   tt.status,
				Header: tt.headers,
			}
			s.StartResponse(resp, rw)
			tt.expect(t, rw)
		})
	}
}
