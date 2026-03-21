package http

import (
	"fmt"
	"net/http"

	"github.com/Rom4eg/gostub/internal/stub"
)

func (s *Service) HandlerApi(w http.ResponseWriter, r *http.Request) {
	s.l.Debug("Enter HandlerApi")
	defer s.l.Debug("Exit HandlerApi")

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

	resp := &Response{
		Code:   ctx.Code(),
		Body:   body,
		Header: ctx.Headers(),
	}

	s.StartResponse(resp, w)
}
