package pool

import (
	"sync"
)

type Object struct {
	Name string
}

var objectPool = sync.Pool{
	New: func() interface{} {
		return new(Object)
	},
}

func New() *Object {
	return objectPool.Get().(*Object)
}

func (o *Object) Put() {
	objectPool.Put(o)
}
