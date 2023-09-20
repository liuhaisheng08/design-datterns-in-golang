package factory

import (
	"testing"
)

func TestNewProduct(t *testing.T) {
	srv := NewServer()
	t.Log(srv)
}
