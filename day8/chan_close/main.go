package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 10)

	for i := 1; i <= 10; i++ {
		intChan <- i
	}
	close(intChan)

	go func() {
		for ; ; {
			b, ok := <-intChan
			if ok == false {
				fmt.Println("channel is close")
				break
			}
			fmt.Println(b)

		}
	}()

	time.Sleep(time.Second)

}
