package context

func (c *StubContext) AddHeader(k, v string) (string, error) {
	c.headers.Add(k, v)
	return "", nil
}
