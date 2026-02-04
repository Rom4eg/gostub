package template

func (tx *TemplateContext) SetStatus(status int) (string, error) {
	tx.status = status
	return "", nil
}
