package template

import "fmt"

func (tx *TemplateContext) GetStatus() (string, error) {
	return fmt.Sprintf("%d", tx.status), nil
}
