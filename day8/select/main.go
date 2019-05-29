package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 100)
	ch2:=make(chan int,100)
	for i := 1; i <= 100; i++ {
		ch1 <- i
		ch2<-i*i
	}
	for {
		select {
		case <-ch1:
			fmt.Println(" ch1 succ")
			time.Sleep(time.Second)
		case <-ch2:
			fmt.Println(" ch2 succ")
			time.Sleep(time.Second)
		default:
			fmt.Println("zoule")
		}
	}
}
