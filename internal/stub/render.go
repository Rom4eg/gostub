package stub

import (
	"bytes"
	"text/template"
)

var funcs = template.FuncMap{
	"base64encode": base64Encode,
	"base64decode": base64Decode,
	"env":          env,
}

func (s *Service) Render(t string) ([]byte, error) {
	buf, e := s._render(t)
	if e != nil {
		return nil, e
	}
	return buf.Bytes(), nil
}

func (s *Service) RenderString(t string) (string, error) {
	buf, e := s._render(t)
	if e != nil {
		return "", e
	}
	return buf.String(), nil
}

func (s *Service) _render(t string) (*bytes.Buffer, error) {
	files, err := s.Lookup(t)
	if err != nil {
		return nil, err
	}

	tpl, err := template.New("main").Funcs(funcs).ParseFiles(files...)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = tpl.ExecuteTemplate(&buf, "main", s.ctx)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}
