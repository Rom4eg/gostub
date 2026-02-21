package http

import (
	"gostub/internal/stub"
	"net/http"
)

func (s *Service) Handler(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(r)
	ss := stub.New(s.Root, ctx)
	body, err := ss.Render(r.URL.EscapedPath())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(ctx.Code())
	_, err = w.Write(body)
	if err != nil {
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
