package http

func (s *Service) Start() error {
	s.l.Debug("Enter Start")
	defer s.l.Debug("Exit Start")

	return s.srv.ListenAndServe()
}
