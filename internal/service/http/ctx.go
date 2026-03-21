package http

import (
	"encoding/json"
	"net/http"
	"strings"
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

func (c *StubContext) SetHeader(k, v string) (string, error) {
	c.headers.Set(k, v)
	return "", nil
}

func (c *StubContext) DeleteHeader(k string) (string, error) {
	c.headers.Del(k)
	return "", nil
}

func (c *StubContext) JsonBody() (map[string]any, error) {
	b := make(map[string]any)
	if c.Request.Body == nil {
		return b, nil
	}
	defer c.Request.Body.Close()

	if err := json.NewDecoder(c.Request.Body).Decode(&b); err != nil {
		return nil, err
	}
	return b, nil
}

func (c *StubContext) FormBody() (map[string]any, error) {
	b := make(map[string]any)
	if c.Request.Body == nil {
		return b, nil
	}
	defer c.Request.Body.Close()

	if err := c.Request.ParseForm(); err != nil {
		return nil, err
	}
	for k, v := range c.Request.Form {
		b[k] = v
	}
	return b, nil
}

func (c *StubContext) MultipartBody() (map[string]any, error) {
	b := make(map[string]any)
	if c.Request.Body == nil {
		return b, nil
	}
	defer c.Request.Body.Close()

	if err := c.Request.ParseMultipartForm(0); err != nil {
		return nil, err
	}
	for k, v := range c.Request.MultipartForm.Value {
		b[k] = v
	}
	return b, nil
}

func (c *StubContext) Body() (map[string]any, error) {
	if c.Request.Body == nil {
		return make(map[string]any), nil
	}
	defer c.Request.Body.Close()

	switch ct := strings.ToLower(c.Request.Header.Get("content-type")); {
	case ct == "application/json":
		return c.JsonBody()
	case ct == "application/x-www-form-urlencoded":
		return c.FormBody()
	case strings.Contains(ct, "multipart/form-data"):
		return c.MultipartBody()
	}

	return make(map[string]any), nil
}
