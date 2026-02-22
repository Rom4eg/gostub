package stub

import (
	"encoding/base64"
	"os"
)

func base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func base64Decode(data string) []byte {
	dec, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return []byte(err.Error())
	}
	return dec
}

func env(name string) string {
	return os.Getenv(name)
}
