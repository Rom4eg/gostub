package http

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	lm "github.com/Rom4eg/gostub/log/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_Handler(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		rw      *httptest.ResponseRecorder
		r       *http.Request
		expect  func(*testing.T, *httptest.ResponseRecorder)
		prepare func(*testing.T, *lm.MockILogger)
		target  string
		tpl     string
	}{
		{
			name:   "FAIL: Render error",
			rw:     httptest.NewRecorder(),
			r:      httptest.NewRequest(http.MethodGet, "/api", nil),
			target: "foo",
			tpl:    "",
			expect: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, w.Code)
				assert.NotNil(t, w.Body)
			},
			prepare: func(t *testing.T, mi *lm.MockILogger) {
				mi.EXPECT().Info(gomock.Any()).AnyTimes()
				mi.EXPECT().Debug(gomock.Any()).AnyTimes()
				mi.EXPECT().Error(gomock.Any()).AnyTimes()
			},
		},
		{
			name:   "PASS: No response code and content not empty",
			rw:     httptest.NewRecorder(),
			r:      httptest.NewRequest(http.MethodGet, "/api", nil),
			target: "api",
			tpl:    "{{define \"main\"}}PASS{{- end -}}",
			expect: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, w.Code)
				assert.NotNil(t, w.Body)
			},
			prepare: func(t *testing.T, mi *lm.MockILogger) {
				mi.EXPECT().Info(gomock.Any()).AnyTimes()
				mi.EXPECT().Debug(gomock.Any()).AnyTimes()
				mi.EXPECT().Error(gomock.Any()).AnyTimes()
			},
		},
		{
			name:   "PASS: No response code and content is empty",
			rw:     httptest.NewRecorder(),
			r:      httptest.NewRequest(http.MethodGet, "/api", nil),
			target: "api",
			tpl:    "{{- define \"main\" -}}{{- end -}}",
			expect: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotImplemented, w.Code)
				assert.NotNil(t, w.Body)
			},
			prepare: func(t *testing.T, mi *lm.MockILogger) {
				mi.EXPECT().Info(gomock.Any()).AnyTimes()
				mi.EXPECT().Debug(gomock.Any()).AnyTimes()
				mi.EXPECT().Error(gomock.Any()).AnyTimes()
			},
		},
		{
			name:   "PASS: OK",
			rw:     httptest.NewRecorder(),
			r:      httptest.NewRequest(http.MethodGet, "/api", nil),
			target: "api",
			tpl:    "{{- define \"main\" -}}{{- .SetCode 201 -}}Some content{{- end -}}",
			expect: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, w.Code)
				assert.Equal(t, "Some content", w.Body.String())
			},
			prepare: func(t *testing.T, mi *lm.MockILogger) {
				mi.EXPECT().Info(gomock.Any()).AnyTimes()
				mi.EXPECT().Debug(gomock.Any()).AnyTimes()
				mi.EXPECT().Error(gomock.Any()).AnyTimes()
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

			root := t.TempDir()
			serviceName := "test"
			targetPath := filepath.Join(root, serviceName, tt.target)
			err := os.MkdirAll(targetPath, os.ModePerm)
			assert.NoError(t, err)

			fh, err := os.Create(filepath.Join(targetPath, "index"))
			assert.NoError(t, err)
			defer fh.Close()

			_, err = fh.WriteString(tt.tpl)
			assert.NoError(t, err)

			opts := ServiceOpts{
				Root: root,
			}

			s := New("test", l, opts)
			s.Handler(tt.rw, tt.r)
			tt.expect(t, tt.rw)
		})
	}
}
