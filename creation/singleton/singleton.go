package singleton

import (
	"sync"

	"github.com/haishengliu/design-datterns-in-golang/creation"
)

type Singleton interface{}

var (
	instance *creation.Server
	once     sync.Once
)

func New() Singleton {
	once.Do(func() {
		instance = &creation.Server{
			// do something here
		}
	})
	return instance
}
