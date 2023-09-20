package factory

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	srv, _ := NewServer("http")
	t.Log(srv.Protocol)
}
