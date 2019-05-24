package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var lock sync.Mutex
	map1 := make(map[int]int)
	map1[2] = 22
	map1[3] = 33
	map1[4] = 44
	map1[5] = 55
	for i := 1; i <= 100; i++ {
		go func() {
			lock.Lock()
			map1[2] = rand.Intn(100)

			lock.Unlock()

		}()
	}
	time.Sleep(time.Second)

	lock.Lock()
	fmt.Println(map1)
	lock.Unlock()

}
