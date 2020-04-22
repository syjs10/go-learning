package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var rwlock sync.RWMutex

func testMap() {
	var a map[int]int
	var count int32
	a = make(map[int]int, 5)
	a[8] = 10
	a[3] = 10
	a[1] = 10
	a[2] = 10
	a[18] = 10

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			rwlock.Lock()
			// b[1] = rand.Intn(100)
			time.Sleep(10 * time.Microsecond)
			rwlock.Unlock()
		}(a)
	}
	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			for {
				rwlock.RLock()
				// fmt.Println(a)
				time.Sleep(time.Microsecond)
				rwlock.RUnlock()
				atomic.AddInt32(&count, 1)
			}

		}(a)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count))
}

func main() {
	testMap()
}
