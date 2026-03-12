package http

import (
	"fmt"
	"net/http"
)

func (s *Service) Start() error {
	s.l.Debug("Enter Start")
	defer s.l.Debug("Exit Start")

	if http, ok := s.srv.(*http.Server); ok {
		s.l.Info(fmt.Sprintf("listen on %s", http.Addr))
	}

	return s.srv.ListenAndServe()
}
