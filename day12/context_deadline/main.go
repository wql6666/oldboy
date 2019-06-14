package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*4))

	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println("jieshu")
	case <-time.After(time.Second * 5):
		fmt.Println("overslept")
	}

}
