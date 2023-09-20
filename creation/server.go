package creation

import (
	"crypto/tls"
	"time"
)

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConn  int
	TLS      *tls.Config
}
