package builder

import (
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	b := Builder{}
	b.CreateServer("localhost", 3306).
		WithProtocol("tcp").
		WithMaxConn(10).Build()
}
