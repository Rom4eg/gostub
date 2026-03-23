package context

import (
	"net/http"
)

type StubContext struct {
	Request *http.Request

	code    int
	headers http.Header
}

func New(r *http.Request) *StubContext {
	return &StubContext{
		Request: r,
		headers: make(http.Header),
	}
}
