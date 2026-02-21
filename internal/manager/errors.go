package manager

import "fmt"

var (
	ErrServiceAlreadyStarted = fmt.Errorf("service already started")
	ErrServiceNotStarted     = fmt.Errorf("service not started")
)
