package service

import (
	"fmt"
	"net/http"

	httpsrv "github.com/Rom4eg/gostub/internal/service/http"
	"github.com/Rom4eg/gostub/log"
)

type Service interface {
	Start() error
	Stop() error
}

type FactoryOpt struct {
	Type       ServiceType
	Logger     log.ILogger
	ServiceOpt map[string]any
}

type Factory struct {
}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) MakeService(name string, o FactoryOpt) (Service, error) {
	switch o.Type {
	case ServiceHttp, ServiceApi:
		if o.Type == ServiceHttp {
			warn := fmt.Sprintf("Service type \"%s\" is deprecated and will be removed in the future. Use \"%s\" instead", ServiceHttp, ServiceApi)
			o.Logger.Warning(warn)
		}

		opts, err := httpsrv.NewServiceOpts(o.ServiceOpt)
		if err != nil {
			return nil, err
		}
		srv := httpsrv.New(name, o.Logger, opts)
		srv.HandlerFunc = http.HandlerFunc(srv.HandlerApi)
		return srv, nil
	default:
		return nil, ErrUnknownServiceType
	}
}
