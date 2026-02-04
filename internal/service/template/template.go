package template

import (
	"net/http"
)

type TemplateContext struct {
	Request  *http.Request
	Response http.ResponseWriter
	status   int
}

func NewTemplateContext(r *http.Request, w http.ResponseWriter) *TemplateContext {
	return &TemplateContext{
		Request:  r,
		Response: w,
		status:   http.StatusNotImplemented,
	}
}
