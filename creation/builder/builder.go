package builder

import (
	"crypto/tls"
	"time"

	"github.com/haishengliu/design-datterns-in-golang/creation"
)

type Builder struct {
	creation.Server
}

func (sb *Builder) CreateServer(addr string, port int) *Builder {
	sb.Server.Addr = addr
	sb.Server.Port = port
	return sb
}

func (sb *Builder) WithProtocol(protocol string) *Builder {
	sb.Server.Protocol = protocol
	return sb
}

func (sb *Builder) WithMaxConn(maxConn int) *Builder {
	sb.Server.MaxConn = maxConn
	return sb
}

func (sb *Builder) WithTimeOut(timeout time.Duration) *Builder {
	sb.Server.Timeout = timeout
	return sb
}

func (sb *Builder) WithTLS(tls *tls.Config) *Builder {
	sb.Server.TLS = tls
	return sb
}

func (sb *Builder) Build() creation.Server {
	return sb.Server
}
