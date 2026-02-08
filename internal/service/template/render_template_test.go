package template

import (
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderTemplate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		file      string
		ctx       *TemplateContext
		tpl       []byte
		wantTpl   []byte
		wantError bool
	}{
		{
			name:      "FAIL: file not exists",
			file:      "not-existing",
			ctx:       &TemplateContext{},
			tpl:       []byte(""),
			wantTpl:   nil,
			wantError: true,
		},
		{
			name:      "FAIL: invalid template",
			file:      "mytpl.tpl",
			ctx:       &TemplateContext{},
			tpl:       []byte(`{{ the error is here }}`),
			wantTpl:   nil,
			wantError: true,
		},
		{
			name: "PASS: render template",
			file: "mytpl.tpl",
			ctx: &TemplateContext{
				Request: &http.Request{
					URL: &url.URL{
						Path: "/my-path",
					},
				},
			},
			tpl:       []byte(`{{- .Request.URL.Path -}}`),
			wantTpl:   []byte("/my-path"),
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dir := t.TempDir()
			tmpFile := filepath.Join(dir, tt.file)
			err := os.WriteFile(tmpFile, tt.tpl, 0777)
			assert.NoError(t, err)

			actual, err := RenderTemplate(tmpFile, tt.ctx)
			if tt.wantError {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.wantTpl, actual)
		})
	}
}
