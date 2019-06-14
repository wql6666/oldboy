package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := test(ctx)
	for b := range c {
		fmt.Println(b)
		if b == 5 {
			break
		}
	}
}

func test(ctx context.Context) <-chan int {
	c := make(chan int)
	n := 1
	go func() {

		for {
			select {
			case <-ctx.Done():
				return
			case c <- n:

				n++
			}
		}
	}()
	return c
}
