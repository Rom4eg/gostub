package context

func (c *StubContext) DeleteHeader(k string) (string, error) {
	c.headers.Del(k)
	return "", nil
}
