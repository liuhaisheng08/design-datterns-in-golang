package option

import (
	"crypto/tls"
	"github.com/haishengliu/design-datterns-in-golang/creation"
	"time"
)

type Option func(server *creation.Server)

func WithProtocol(p string) Option {
	return func(s *creation.Server) {
		s.Protocol = p
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *creation.Server) {
		s.Timeout = timeout
	}
}

func WithMaxConn(maxConn int) Option {
	return func(s *creation.Server) {
		s.MaxConn = maxConn
	}
}

func WithTLS(tls *tls.Config) Option {
	return func(s *creation.Server) {
		s.TLS = tls
	}
}

func NewServer(addr string, port int, options ...func(server *creation.Server)) *creation.Server {
	srv := creation.Server{
		Addr: addr,
		Port: port,
	}
	for _, option := range options {
		option(&srv)
	}
	return &srv
}
