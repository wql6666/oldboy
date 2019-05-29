package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Second)
	t1 := time.After(time.Second * 2)
	for ; ; {
		select {
		case <-t.C:
			fmt.Println("timeout")
		case <-t1:
			fmt.Println("time.after")
		default:
			fmt.Println("lala")
			time.Sleep(time.Second)
		}

		t.Stop()

	}
}
