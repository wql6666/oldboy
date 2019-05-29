package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan int)
	wg.Add(1)
	go func(intCh chan int) {
		for i := 1; i <= 100; i++ {
			ch1 <- i
		}
		close(ch1)
		fmt.Println("send exit")
		wg.Done()
	}(ch1)
	wg.Add(1)
	go func(intCh chan int) {

		for v := range ch1 {
			fmt.Println(v)
		}
		fmt.Println("rec exit")
		wg.Done()
	}(ch1)
	wg.Wait()

}
