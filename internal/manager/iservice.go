package manager

//go:generate mockgen -typed -destination=mocks/manager.go -package=mocks . IService
type IService interface {
	Start() error
	Stop() error
}
