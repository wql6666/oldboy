package main

import (
	"fmt"
)

func main() {
	startChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)

	for i := 1; i <= 1000; i++ {
		startChan <- i
	}
	close(startChan)

	for i := 1; i <= 8; i++ {
		go calc(startChan, resultChan, exitChan)

	}

	for i := 1; i <= 8; i++ {
		<-exitChan
		fmt.Println("wait goroutine", i, "exited")
	}
	close(resultChan)

	for v := range resultChan {
		fmt.Println(v)
	}
	//time.Sleep(time.Second)
}

func calc(startChan, resChan chan int, exitChan chan bool) {
	for v := range startChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			resChan <- v
		}
	}
	//fmt.Println("exit")
	exitChan <- true

}
