package learn

import (
	"sync"
	"sync/atomic"
)

type singleton struct{}

// 单例
var (
	instance    *singleton
	initialized uint32
	mu          sync.Mutex
)

func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}
	return instance
}

// sync.Once 单例
var (
	instance2 *singleton
	once      sync.Once
)

func Instance2() *singleton {
	once.Do(func() {
		instance2 = &singleton{}
	})
	return instance2
}
