package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan struct{},2)
	go func(intCh chan int, StruCh chan struct{}) {
		var stru struct{}
		for i := 1; i <= 100; i++ {
			ch1 <- i
		}
		ch2 <- stru
		fmt.Println("send exit")
		close(ch1)
	}(ch1, ch2)

	go func(intCh chan int, StruCh chan struct{}) {
		var stru struct{}

		for v := range ch1 {
			fmt.Println(v)
		}
		/*for ; ; {
			v, ok := <-ch1
			if !ok {
				fmt.Println("ok----",ok)
				break
			}
			fmt.Println(v)
		} */
		ch2 <- stru
		fmt.Println("rec exit")
	}(ch1, ch2)
	<-ch2
	<-ch2

	//time.Sleep(time.Second)
}
