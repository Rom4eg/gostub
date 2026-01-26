package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReload(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		cfgFile string
		panic   error
	}{
		{
			name:    "FAIL: non existing file",
			cfgFile: "non-existing-file.yaml",
			panic:   os.ErrNotExist,
		},
		{
			name:    "FAIL: empty config",
			cfgFile: "./test_data/empty_test.yaml",
			panic:   ErrEmptyConfig,
		},
		{
			name:    "PASS: OK",
			cfgFile: "./test_data/pass_test.yaml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			SetConfigLocation(tt.cfgFile)
			if tt.panic != nil {
				assert.Panics(t, Reload, tt.panic)
			}
		})
	}
}
