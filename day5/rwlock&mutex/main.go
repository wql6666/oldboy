package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var map1 = map[int]int{
		1: 11,
		2: 22,
		3: 33,
		4: 44,
		5: 55,
	}
	testRwlock(map1)
	testMutexlock(map1)

}

func testMutexlock(b map[int]int) {
	var count int32
	var rwlock sync.Mutex
	//for i := 1; i <= 2; i++ {
	//	go func(map1 map[int]int) {
	//		for ; ; {
	//			rwlock.Lock()
	//			time.Sleep(time.Millisecond * 10)
	//			b[1] = 100
	//			rwlock.Unlock()
	//			atomic.AddInt32(&count, 1)
	//		}
	//
	//	}(b)
	//}
	for i := 1; i <= 100; i++ {

		go func() {
			for ; ; {
				rwlock.Lock()
				time.Sleep(time.Millisecond)
				atomic.AddInt32(&count, 1)
				rwlock.Unlock()
			}

		}()

	}
	time.Sleep(time.Second * 3)

	fmt.Printf("互斥锁的并发数量是：%d", atomic.LoadInt32(&count))
}

func testRwlock(b map[int]int) {
	var count int32
	var rwlock sync.RWMutex
	//for i := 1; i <= 2; i++ {
	//	go func(map1 map[int]int) {
	//		for ; ; {
	//			rwlock.Lock()
	//			time.Sleep(time.Millisecond * 10)
	//			b[1] = 100
	//			rwlock.Unlock()
	//			atomic.AddInt32(&count, 1)
	//		}
	//	}(b)
	//}
	for i := 1; i <= 100; i++ {

		go func() {
			for ; ; {
				rwlock.RLock()
				time.Sleep(time.Millisecond)
				atomic.AddInt32(&count, 1)
				rwlock.RUnlock()
			}
		}()

	}
	time.Sleep(time.Second * 3)

	fmt.Printf("读写锁的并发次数为：%d\n", atomic.LoadInt32(&count))
}
