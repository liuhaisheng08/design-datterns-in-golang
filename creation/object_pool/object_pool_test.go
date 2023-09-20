package pool

import (
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestNewPool(t *testing.T) {
	obj := New()
	defer obj.Put()
}
