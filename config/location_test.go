package config

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGetSetConfigLocation(t *testing.T) {
	t.Parallel()

	cfg := "./some+-non_existing=path!.yaml"
	SetConfigLocation(cfg)
	assert.Equal(t, GetConfigLocation(), cfg)
}
