package context

import "strings"

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
