package factory

import "github.com/haishengliu/design-datterns-in-golang/creation"

type Server interface {
}

func NewServer() Server {
	return &creation.Server{
		// do something here
	}
}
