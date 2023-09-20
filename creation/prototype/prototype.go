package prototype

import "github.com/haishengliu/design-datterns-in-golang/creation"

type Cloneable interface {
	Clone() Cloneable
}

type Server struct {
	creation.Server
}

func NewServer() Cloneable {
	return &Server{
		// do something here
	}
}

func (e *Server) Clone() Cloneable {
	res := *e
	return &res
}
