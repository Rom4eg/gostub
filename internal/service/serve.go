package service

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path"
	"text/template"
)

func (s *Service) Serve(w http.ResponseWriter, r *http.Request) {
	s.l.Info(fmt.Sprintf("Serve %s", r.URL.Path))

	file := path.Join(s.root, r.URL.Path)
	if _, err := os.Stat(file); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	ctx := &TemplateContext{
		Request:  r,
		Response: w,
	}

	// fh, err := os.Open(file)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// defer fh.Close()

	// var txtTemplate
	// tplText := io.Copy()

	t, err := template.ParseFiles(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	e := t.Execute(&buf, ctx)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(ctx.status)
	_, err = w.Write(buf.Bytes())
	if err != nil {
		s.l.Error(err.Error())
		return
	}
}
