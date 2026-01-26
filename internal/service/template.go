package service

import (
	"bytes"
	"net/http"
	"text/template"
)

type TemplateContext struct {
	Request  *http.Request
	Response http.ResponseWriter
	status   int
}

func (tx *TemplateContext) AddHeader(name string, value string) (string, error) {
	tx.Response.Header().Add(name, value)
	return "", nil
}

func (tx *TemplateContext) SetStatus(status int) (string, error) {
	tx.status = status
	return "", nil
}

func RenderTemplate(r *http.Request) (string, error) {
	ctx := &TemplateContext{Request: r}
	t := template.New(".")
	var buf bytes.Buffer
	e := t.Execute(&buf, ctx)
	return buf.String(), e
}
