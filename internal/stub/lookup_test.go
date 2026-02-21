package stub

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Lookup(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		path    string
		prepare func(string) error
		expect  func(*testing.T, []string, error)
	}{
		{
			name: "FAIL: not existing",
			path: "not-existing-path",
			prepare: func(s string) error {
				return nil
			},
			expect: func(t *testing.T, res []string, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "FAIL: is dir and index file not exists",
			path: "test-dir",
			prepare: func(s string) error {
				fp := filepath.Join(s, "test-dir")
				err := os.Mkdir(fp, 0777)
				if err != nil {
					return err
				}

				return nil
			},
			expect: func(t *testing.T, res []string, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "PASS: request to directory returns index file",
			path: "test-dir",
			prepare: func(s string) error {
				fp := filepath.Join(s, "test-dir")
				err := os.Mkdir(fp, 0777)
				if err != nil {
					return err
				}

				idx := filepath.Join(fp, "index")
				_, err = os.Create(idx)
				if err != nil {
					return err
				}
				return nil
			},
			expect: func(t *testing.T, res []string, err error) {
				assert.NoError(t, err)
				assert.Len(t, res, 1)
				for _, r := range res {
					assert.True(t, strings.HasSuffix(r, "index"))
				}
			},
		},
		{
			name: "PASS: request to normal file",
			path: filepath.Join("test-dir", "test-file"),
			prepare: func(s string) error {
				fp := filepath.Join(s, "test-dir")
				err := os.Mkdir(fp, 0777)
				if err != nil {
					return err
				}

				idx := filepath.Join(fp, "test-file")
				_, err = os.Create(idx)
				if err != nil {
					return err
				}
				return nil
			},
			expect: func(t *testing.T, res []string, err error) {
				assert.NoError(t, err)
				assert.Len(t, res, 1)
				for _, r := range res {
					assert.True(t, strings.HasSuffix(r, "test-file"))
				}
			},
		},
		{
			name: "PASS: request to normal file with resources",
			path: filepath.Join("test-dir", "test-file"),
			prepare: func(s string) error {
				fp := filepath.Join(s, "test-dir")
				err := os.Mkdir(fp, 0777)
				if err != nil {
					return err
				}

				target := filepath.Join(fp, "test-file")
				_, err = os.Create(target)
				if err != nil {
					return err
				}

				funcs := filepath.Join(fp, "_funcs")
				_, err = os.Create(funcs)
				if err != nil {
					return err
				}
				return nil
			},
			expect: func(t *testing.T, res []string, err error) {
				assert.NoError(t, err)
				assert.Len(t, res, 2)
				endings := make(map[string]bool)
				for _, r := range res {
					assert.True(t, filepath.IsAbs(r))
					endings[filepath.Base(r)] = true
				}

				assert.True(t, endings["test-file"])
				assert.True(t, endings["_funcs"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tmp := t.TempDir()
			err := tt.prepare(tmp)
			assert.NoError(t, err)

			s := New(tmp, nil)
			res, err := s.Lookup(tt.path)
			tt.expect(t, res, err)
		})
	}
}

func TestService__lookupResources(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		path    string
		prepare func(string) error
		expect  func(*testing.T, []string, error)
	}{
		{
			name: "FAIL: not existing",
			path: "not-existing-path",
			prepare: func(s string) error {
				return nil
			},
			expect: func(t *testing.T, res []string, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "FAIL: not a directory",
			path: "file",
			prepare: func(s string) error {
				fp := filepath.Join(s, "file")
				_, err := os.Create(fp)
				return err
			},
			expect: func(t *testing.T, res []string, err error) {
				assert.ErrorIs(t, err, ErrDirectoryExpected)
			},
		},
		{
			name: "FAIL: inaccessable directory",
			path: "dir",
			prepare: func(s string) error {
				fp := filepath.Join(s, "dir")
				err := os.Mkdir(fp, os.FileMode(0000))
				return err
			},
			expect: func(t *testing.T, res []string, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "PASS: empty directory",
			path: "dir",
			prepare: func(s string) error {
				fp := filepath.Join(s, "dir")
				err := os.Mkdir(fp, os.FileMode(0777))
				return err
			},
			expect: func(t *testing.T, res []string, err error) {
				assert.Empty(t, res)
				assert.NoError(t, err)
			},
		},
		{
			name: "PASS: skip not underscored files",
			path: "dir",
			prepare: func(s string) error {
				fp := filepath.Join(s, "dir")
				err := os.Mkdir(fp, os.FileMode(0777))
				if err != nil {
					return err
				}

				file := filepath.Join(fp, "file")
				_, err = os.Create(file)
				if err != nil {
					return err
				}

				file2 := filepath.Join(fp, "file2")
				_, err = os.Create(file2)
				if err != nil {
					return err
				}

				return err
			},
			expect: func(t *testing.T, res []string, err error) {
				assert.Empty(t, res)
				assert.NoError(t, err)
			},
		},
		{
			name: "PASS: skip subdirectories",
			path: "dir",
			prepare: func(s string) error {
				fp := filepath.Join(s, "dir")
				err := os.Mkdir(fp, os.FileMode(0777))
				if err != nil {
					return err
				}

				file := filepath.Join(fp, "file")
				_, err = os.Create(file)
				if err != nil {
					return err
				}

				fp2 := filepath.Join(fp, "dir2")
				err = os.Mkdir(fp2, os.FileMode(0777))
				if err != nil {
					return err
				}

				file2 := filepath.Join(fp2, "file2")
				_, err = os.Create(file2)
				if err != nil {
					return err
				}

				return err
			},
			expect: func(t *testing.T, res []string, err error) {
				assert.Empty(t, res)
				assert.NoError(t, err)
			},
		},
		{
			name: "PASS: mixed files with and without underscore",
			path: "dir",
			prepare: func(s string) error {
				fp := filepath.Join(s, "dir")
				err := os.Mkdir(fp, os.FileMode(0777))
				if err != nil {
					return err
				}

				file := filepath.Join(fp, "file")
				_, err = os.Create(file)
				if err != nil {
					return err
				}

				fp2 := filepath.Join(fp, "dir2")
				err = os.Mkdir(fp2, os.FileMode(0777))
				if err != nil {
					return err
				}

				file2 := filepath.Join(fp2, "file2")
				_, err = os.Create(file2)
				if err != nil {
					return err
				}

				funcs := filepath.Join(fp, "_funcs.tpl")
				_, err = os.Create(funcs)
				if err != nil {
					return err
				}

				return err
			},
			expect: func(t *testing.T, res []string, err error) {
				assert.Len(t, res, 1)
				assert.NoError(t, err)
				endings := make(map[string]bool)
				for _, path := range res {
					assert.True(t, filepath.IsAbs(path))
					endings[filepath.Base(path)] = true
				}
				assert.True(t, endings["_funcs.tpl"])
				assert.False(t, endings["file2"])
				assert.False(t, endings["file"])
				assert.False(t, endings["dir2"])
				assert.False(t, endings["."])
				assert.False(t, endings[".."])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tmp := t.TempDir()

			err := tt.prepare(tmp)
			assert.NoError(t, err)

			s := New(tmp, nil)
			fullPath := filepath.Join(tmp, tt.path)
			res, err := s._lookupResources(fullPath)
			tt.expect(t, res, err)
		})
	}
}
