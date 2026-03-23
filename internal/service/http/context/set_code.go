package context

func (c *StubContext) SetCode(code int) (string, error) {
	c.code = code
	return "", nil
}
