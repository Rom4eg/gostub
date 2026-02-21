package manager

func (m *Manager) StopService(name string) error {
	m.l.Debug("Enter StopService")
	defer m.l.Debug("Exit StopService")

	svc, ok := m.svc[name]
	if !ok {
		return ErrServiceNotStarted
	}
	delete(m.svc, name)

	return svc.Stop()
}
