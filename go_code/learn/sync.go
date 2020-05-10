package learn

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 互斥锁
var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

// 原子操作
var total2 uint64

func worker2(wg *sync.WaitGroup) {
	defer wg.Done()
	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&total2, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker2(&wg)
	go worker2(&wg)
	wg.Wait()
	fmt.Println(total2)
}
