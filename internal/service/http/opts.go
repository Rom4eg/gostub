package http

type ServiceOpts struct {
	Host   string
	Port   int
	Root   string
	Server IServer
}

func NewServiceOpts(opts map[string]any) (ServiceOpts, error) {
	h, ok := opts["host"].(string)
	if !ok {
		return ServiceOpts{}, ErrIncorrectServiceOptions
	}

	p, ok := opts["port"].(int)
	if !ok {
		return ServiceOpts{}, ErrIncorrectServiceOptions
	}

	r, ok := opts["root"].(string)
	if !ok {
		return ServiceOpts{}, ErrIncorrectServiceOptions
	}

	return ServiceOpts{
		Host:   h,
		Port:   p,
		Root:   r,
		Server: nil,
	}, nil
}
