package main

import (
	"fmt"
	"time"
)

const (
	Man    = 1
	Female = 2
)

func main() {
	consttest()
}
func consttest() {
	time := time.Now().Second()
	if time%Female == 0 {
		fmt.Println("female")
	} else {
		fmt.Println("male")
	}
}
