package service

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Url2File(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		path    string
		prepare func(string) error
		expect  func(*testing.T, string, error)
	}{
		{
			name: "FAIL: not existing",
			path: "not-existing",
			prepare: func(s string) error {
				return nil
			},
			expect: func(t *testing.T, s string, err error) {
				assert.Equal(t, "", s)
				assert.Error(t, err)
				assert.True(t, os.IsNotExist(err))
			},
		},
		{
			name: "Fail: path is directory, index file not exists",
			path: "direntry",
			prepare: func(s string) error {
				dir := filepath.Join(s, "direntry")
				err := os.Mkdir(dir, 0755)
				if err != nil {
					return err
				}

				return nil
			},
			expect: func(t *testing.T, s string, err error) {
				assert.Equal(t, "", s)
				assert.Error(t, err)
				assert.True(t, os.IsNotExist(err))
			},
		},
		{
			name: "PASS: path is directory, index exists",
			path: "direntry",
			prepare: func(s string) error {
				dir := filepath.Join(s, "direntry")
				err := os.Mkdir(dir, 0755)
				if err != nil {
					return err
				}

				idx := filepath.Join(dir, "index")
				_, err = os.Create(idx)
				if err != nil {
					return err
				}

				return nil
			},

			expect: func(t *testing.T, s string, err error) {
				assert.NoError(t, err)
				base := filepath.Base(s)
				assert.Equal(t, "index", base)
			},
		},
		{
			name: "PASS: path is file",
			path: "file",
			prepare: func(s string) error {
				file := filepath.Join(s, "file")
				_, err := os.Create(file)
				if err != nil {
					return err
				}

				return nil
			},
			expect: func(t *testing.T, s string, err error) {
				base := filepath.Base(s)
				assert.Equal(t, "file", base)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tmp := t.TempDir()
			tmp_name := filepath.Join(tmp, "test")
			err := os.Mkdir(tmp_name, 0755)
			assert.NoError(t, err)

			err = tt.prepare(tmp_name)
			assert.NoError(t, err)
			opts := ServiceOpts{
				StubRoot: tmp,
				Name:     "test",
			}

			svc := NewService(opts)
			path := filepath.Join(svc.root, tt.path)
			got, err := svc.Url2File(path)
			tt.expect(t, got, err)
		})
	}
}
