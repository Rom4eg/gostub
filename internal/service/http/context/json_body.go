package context

import "encoding/json"

func (c *StubContext) JsonBody() (map[string]any, error) {
	b := make(map[string]any)
	if c.Request.Body == nil {
		return b, nil
	}
	defer c.Request.Body.Close()

	if err := json.NewDecoder(c.Request.Body).Decode(&b); err != nil {
		return b, err
	}
	return b, nil
}
