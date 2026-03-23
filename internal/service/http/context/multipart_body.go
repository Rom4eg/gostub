package context

func (c *StubContext) MultipartBody() (map[string]any, error) {
	b := make(map[string]any)
	if c.Request.Body == nil {
		return b, nil
	}
	defer c.Request.Body.Close()

	if err := c.Request.ParseMultipartForm(0); err != nil {
		return b, err
	}
	for k, v := range c.Request.MultipartForm.Value {
		b[k] = v
	}
	return b, nil
}
