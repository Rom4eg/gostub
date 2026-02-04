package template

import (
	"bytes"
	"text/template"
)

func RenderTemplate(f string, ctx *TemplateContext) ([]byte, error) {
	t, err := template.ParseFiles(f)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, ctx)
	if err != nil {
		return nil, err
	}
	return bytes.Trim(buf.Bytes(), "\n\r\t"), nil
}
