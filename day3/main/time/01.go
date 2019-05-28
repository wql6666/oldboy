package main

import (
	"fmt"
	"time"
)

func main() {
	s := time.Now().Format("2006/01/02 15:04:05")
	fmt.Println(s)
	start := time.Now().UnixNano()
	exam()
	end := time.Now().UnixNano()
	fmt.Println((end - start) * 1000)
	a := 5
	test(&a)
	fmt.Println(a)
	if 1==1{
		fmt.Println(a)
	}else if true {
		fmt.Println(1)
	}
}
func exam() {
	time.Sleep(time.Microsecond * 100)
}

func test(a *int) {
	*a = 100
}
