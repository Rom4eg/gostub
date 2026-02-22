package http

import (
	"fmt"
	"gostub/internal/stub"
	"net/http"
)

func (s *Service) Handler(w http.ResponseWriter, r *http.Request) {
	s.l.Debug("Enter Handler")
	defer s.l.Debug("Exit Handler")

	ctx := NewContext(r)
	ss := stub.New(s.Root, ctx)

	path := r.URL.EscapedPath()
	s.l.Info(fmt.Sprintf("Rendering %s", path))
	body, err := ss.Render(path)
	if err != nil {
		s.l.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(ctx.Code())
	_, err = w.Write(body)
	if err != nil {
		s.l.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	for k, v := range ctx.Headers() {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
}
