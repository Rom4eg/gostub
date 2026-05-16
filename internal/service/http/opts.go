package http

import (
	"net/http"

	"github.com/Rom4eg/map2struct"
)

type ServiceOpts struct {
	Host string `m2s:"host"`
	Port int    `m2s:"port"`
	Root string `m2s:"root"`
	Trim bool   `m2s:"trim"`

	Server  IServer
	Handler http.HandlerFunc
}

func NewServiceOpts(opts map[string]any) (ServiceOpts, error) {
	var o ServiceOpts
	e := map2struct.Map2Struct(opts, &o)
	if e != nil {
		return ServiceOpts{}, e
	}

	if o.Host == "" || o.Port == 0 || o.Root == "" {
		return ServiceOpts{}, ErrIncorrectServiceOptions
	}

	_, ok := opts["trim"]
	if !ok {
		o.Trim = true
	}
	return o, nil
}
