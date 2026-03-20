package http

import "net/http"

func (s *Service) Handler(w http.ResponseWriter, r *http.Request) {
	s.l.Debug("Enter Handler")
	defer s.l.Debug("Exit Handler")

	if s.HandlerFunc != nil {
		s.HandlerFunc(w, r)
		return
	}

	resp := &Response{}
	s.StartResponse(resp, w)
}
