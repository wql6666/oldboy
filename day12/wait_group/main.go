package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		//	go func(i int) {
		//		fmt.Println(i)
		//		wg.Done()
		//	}(i)
		//}
		go print(&wg, i)
	}
	wg.Wait()

}

func print(wg *sync.WaitGroup, i int) {
	fmt.Println(i)
	wg.Done()
}