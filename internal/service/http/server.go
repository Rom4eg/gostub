package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Rom4eg/gostub/log"
)

//go:generate mockgen -typed -destination=mocks/server.go -package=mocks . IServer

type IServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

func DefaultServer(host string, port int, handler http.Handler, l log.ILogger) IServer {
	addr := fmt.Sprintf("%s:%d", host, port)
	l.Info(fmt.Sprintf("Service configured to listen %s", addr))
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}
