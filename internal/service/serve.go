package service

import (
	"fmt"
	"gostub/internal/service/template"
	"net/http"
	"os"
	"path/filepath"
)

func (s *Service) Serve(w http.ResponseWriter, r *http.Request) {
	s.l.Info(fmt.Sprintf("Serve %s", r.URL.Path))

	path := filepath.Join(s.root, r.URL.Path)
	file, err := s.Url2File(path)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := template.NewTemplateContext(r, w)
	res, err := template.RenderTemplate(file, ctx)
	if err != nil {
		s.l.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(ctx.GetStatusCode())
	_, err = w.Write(res)
	if err != nil {
		s.l.Error(err.Error())
		return
	}
}
