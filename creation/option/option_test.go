package option

import (
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	srv := NewServer(
		"localhost", 3306,
		WithProtocol("tcp"), WithTimeout(1*time.Second),
	)
	t.Log(srv)
}
