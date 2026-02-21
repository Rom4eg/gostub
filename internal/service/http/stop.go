package http

import (
	"context"
	"time"
)

func (s *Service) Stop() error {
	s.l.Debug("Enter Stop")
	defer s.l.Debug("Exit Stop")

	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	return s.srv.Shutdown(ctx)
}
