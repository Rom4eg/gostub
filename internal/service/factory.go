package service

import (
	httpsrv "gostub/internal/service/http"
	"gostub/log"
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
	case ServiceHttp:
		opts, err := httpsrv.NewServiceOpts(o.ServiceOpt)
		if err != nil {
			return nil, err
		}
		return httpsrv.New(name, o.Logger, opts), nil
	default:
		return nil, ErrUnknownServiceType
	}
}
