package context

func (c *StubContext) SetHeader(k, v string) (string, error) {
	c.headers.Set(k, v)
	return "", nil
}
