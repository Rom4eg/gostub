package context

import "net/http"

func (c *StubContext) Headers() http.Header {
	return c.headers
}
