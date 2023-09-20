package factory

import (
	"errors"

	"github.com/haishengliu/design-datterns-in-golang/creation"
)

type HttpServer struct {
	creation.Server
}

type RPCServer struct {
	creation.Server
}

func NewServer(t string) (*creation.Server, error) {
	switch t {
	case "http":
		return &creation.Server{
			Protocol: "http",
		}, nil
	case "rpc":
		return &creation.Server{
			Protocol: "rpc",
		}, nil
	default:
		return nil, errors.New("err")
	}
}
