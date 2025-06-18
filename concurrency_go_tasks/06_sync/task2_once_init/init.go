package initonce

import (
	"sync"
	"time"
)

var (
	once        sync.Once
	initialized bool
)

func Init() {
	once.Do(func() {
		time.Sleep(10 * time.Millisecond)
		initialized = true
	})
}

func Initialized() bool {
	Init()
	return initialized
}
