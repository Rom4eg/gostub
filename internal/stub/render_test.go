package stub

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService__render(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		file    string
		tpl     string
		prepare func(string, string) error
		expect  func(*testing.T, *bytes.Buffer, error)
	}{
		{
			name: "FAIL: not existing",
			file: "not-existing",
			tpl:  "{{define \"main\"}}{{end}}",
			prepare: func(s, tpl string) error {
				return nil
			},
			expect: func(t *testing.T, buf *bytes.Buffer, err error) {
				assert.Nil(t, buf)
				assert.Error(t, err)
			},
		},
		{
			name: "FAIL: invalid template",
			file: "index",
			tpl:  "{{ invalid template }}",
			prepare: func(s, tpl string) error {
				fp := filepath.Join(s, "index")
				fh, err := os.OpenFile(fp, os.O_CREATE|os.O_RDWR, 0777)
				if err != nil {
					return err
				}
				defer fh.Close()

				_, err = fh.WriteString(tpl)
				return err
			},
			expect: func(t *testing.T, buf *bytes.Buffer, err error) {
				assert.Nil(t, buf)
				assert.Error(t, err)
			},
		},
		{
			name: "PASS: valid template",
			file: "index",
			tpl:  "{{define \"main\"}}{{- if eq 1 1 -}}PASS{{- end -}}{{- end -}}",
			prepare: func(s, tpl string) error {
				fp := filepath.Join(s, "index")
				fh, err := os.OpenFile(fp, os.O_CREATE|os.O_RDWR, 0777)
				if err != nil {
					return err
				}
				defer fh.Close()

				_, err = fh.WriteString(tpl)
				return err
			},
			expect: func(t *testing.T, buf *bytes.Buffer, err error) {
				str := buf.String()
				assert.Equal(t, "PASS", str)
				assert.NoError(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tmp := t.TempDir()
			e := tt.prepare(tmp, tt.tpl)
			assert.NoError(t, e)

			s := New(tmp, nil)
			buf, e := s._render(tt.file)
			tt.expect(t, buf, e)
		})
	}
}

func TestService_RenderString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		tpl     string
		prepare func(string, string) error
	}{
		{
			name: "PASS: render string",
			tpl:  "{{define \"main\"}}{{- if eq 1 1 -}}PASS{{- end -}}{{- end -}}",
			prepare: func(s, tpl string) error {
				fp := filepath.Join(s, "index")
				fh, err := os.OpenFile(fp, os.O_CREATE|os.O_RDWR, 0777)
				if err != nil {
					return err
				}
				defer fh.Close()

				_, err = fh.WriteString(tpl)
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tmp := t.TempDir()
			e := tt.prepare(tmp, tt.tpl)
			assert.NoError(t, e)
			s := New(tmp, nil)
			str, e := s.RenderString("index")
			assert.NoError(t, e)
			assert.Equal(t, "PASS", str)
		})
	}
}

func TestService_Render(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		tpl     string
		prepare func(string, string) error
	}{
		{
			name: "PASS: render string",
			tpl:  "{{define \"main\"}}{{- if eq 1 1 -}}PASS{{- end -}}{{- end -}}",
			prepare: func(s, tpl string) error {
				fp := filepath.Join(s, "index")
				fh, err := os.OpenFile(fp, os.O_CREATE|os.O_RDWR, 0777)
				if err != nil {
					return err
				}
				defer fh.Close()

				_, err = fh.WriteString(tpl)
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tmp := t.TempDir()
			e := tt.prepare(tmp, tt.tpl)
			assert.NoError(t, e)
			s := New(tmp, nil)
			b, e := s.Render("index")
			assert.NoError(t, e)
			assert.Equal(t, []byte("PASS"), b)
		})
	}
}
