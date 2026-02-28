package http

import (
	"net/http"
	"path/filepath"

	"github.com/Rom4eg/gostub/log"
)

type Service struct {
	Host string
	Port int
	Root string

	l   log.ILogger
	srv IServer
}

func New(name string, l log.ILogger, opts ServiceOpts) *Service {
	svc := &Service{
		Host: opts.Host,
		Port: opts.Port,
		l:    l,
		srv:  opts.Server,
		Root: filepath.Join(opts.Root, name),
	}

	if svc.srv == nil {
		svc.srv = DefaultServer(svc.Host, svc.Port, http.HandlerFunc(svc.Handler), l)
	}

	return svc
}
