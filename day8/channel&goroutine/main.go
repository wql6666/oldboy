package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go write(ch1)
	for {
		go read(ch1)
		time.Sleep(time.Microsecond*100)

	}
	time.Sleep(time.Second)
}

func write(ch1 chan int) {
	for i := 1; i <= 100; i++ {
		ch1 <- i
	}

}

func read(ch1 chan int) {
	s := <-ch1
	fmt.Println(s)
}
