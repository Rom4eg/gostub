package stub

import "encoding/base64"

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
