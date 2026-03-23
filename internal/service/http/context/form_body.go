package context

func (c *StubContext) FormBody() (map[string]any, error) {
	b := make(map[string]any)
	if c.Request.Body == nil {
		return b, nil
	}
	defer c.Request.Body.Close()

	if err := c.Request.ParseForm(); err != nil {
		return b, err
	}
	for k, v := range c.Request.Form {
		b[k] = v
	}
	return b, nil
}
