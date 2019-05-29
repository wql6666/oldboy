package main

import (
	"fmt"
	"sync"
	"time"
)

type task struct {
	n int
}

func main() {

	for i := 1; i <= 20; i++ {
		t := &task{n: i}
	go	test(t)
	}
	time.Sleep(time.Second * 3)
	//lock.Lock()
	//for i, v := range m {
	//	fmt.Printf("%d!=%v\n", i, v)
	//}
	//lock.Unlock()
}

var (
	m    = make(map[int]int)
	lock sync.Mutex
)

func test(t *task) {
	var sum int = 1
	for i := 1; i <= t.n; i++ {
		sum *= i

	}
	fmt.Println(t.n,sum)
	lock.Lock()
	m[t.n] = sum
	lock.Unlock()
}
