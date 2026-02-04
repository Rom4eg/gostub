package template

import "encoding/base64"

func (tx *TemplateContext) Base64Encode(str string) (string, error) {
	e := base64.StdEncoding.EncodeToString([]byte(str))
	return e, nil
}
