package service

import (
	"context"
	"fmt"
	"gostub/log"
	"net/http"
	"path"
)

type ServiceOpts struct {
	Host     string
	Port     string
	Name     string
	StubRoot string
}

type Service struct {
	host string
	port string
	l    log.ILogger
	root string
}

func NewService(opts ServiceOpts) *Service {
	return &Service{
		host: opts.Host,
		port: opts.Port,
		l:    log.With(fmt.Sprintf("[%s]", opts.Name)),
		root: path.Join(opts.StubRoot, opts.Name),
	}
}

func (s *Service) Run(ctx context.Context) error {
	s.l.Info("Run service")
	defer s.l.Info("Exit service")

	addr := fmt.Sprintf("%s:%s", s.host, s.port)
	srv := &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(s.Serve),
	}

	go func(l log.ILogger) {
		<-ctx.Done()
		e := srv.Shutdown(ctx)
		if e != nil {
			l.Error(e.Error())
			return
		}
	}(s.l)

	s.l.Info(fmt.Sprintf("Listen on %s", addr))
	err := srv.ListenAndServe()
	if err != nil {
		s.l.Error(err.Error())
		return err
	}
	return nil
}
