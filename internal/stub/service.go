package stub

type Service struct {
	root string
	ctx  any
}

func New(root string, ctx any) *Service {
	return &Service{
		root: root,
		ctx:  ctx,
	}
}
