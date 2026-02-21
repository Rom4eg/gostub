package stub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64Encode(t *testing.T) {
	assert.Equal(t, base64Encode("test"), "dGVzdA==")
}

func TestBase64Decode(t *testing.T) {
	assert.Equal(t, base64Decode("dGVzdA=="), []byte("test"))
}
