package manager

import "github.com/Rom4eg/gostub/log"

type IManager interface {
	StartService(name string, s IService) error
	StopService(name string) error
}

type Manager struct {
	l   log.ILogger
	svc map[string]IService
}

func New(l log.ILogger) IManager {
	l.Debug("Enter manager constructor")
	defer l.Debug("Exit manager constructor")

	m := &Manager{}
	m.l = l
	m.svc = make(map[string]IService)
	return m
}
