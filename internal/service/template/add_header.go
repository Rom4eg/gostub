package template

func (tx *TemplateContext) AddHeader(name string, value string) (string, error) {
	tx.Response.Header().Add(name, value)
	return "", nil
}
