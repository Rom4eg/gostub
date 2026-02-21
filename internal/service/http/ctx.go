package http

import (
	"net/http"
)

type StubContext struct {
	Request *http.Request

	code    int
	headers http.Header
}

func NewContext(r *http.Request) *StubContext {
	return &StubContext{
		Request: r,
		headers: make(http.Header),
	}
}

func (c *StubContext) Code() int {
	return c.code
}

func (c *StubContext) SetCode(code int) (string, error) {
	c.code = code
	return "", nil
}

func (c *StubContext) Headers() http.Header {
	return c.headers
}

func (c *StubContext) AddHeader(k, v string) (string, error) {
	c.headers.Add(k, v)
	return "", nil
}

func (c *StubContext) DeleteHeader(k string) (string, error) {
	c.headers.Del(k)
	return "", nil
}
