package manager

func (m *Manager) StartService(name string, s IService) error {
	m.l.Debug("Enter StartService")
	defer m.l.Debug("Exit StartService")

	if _, ok := m.svc[name]; ok {
		return ErrServiceAlreadyStarted
	}

	m.svc[name] = s
	return s.Start()
}
