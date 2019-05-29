package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	stu1 := student{
		"alan",
		18,
	}
	var ch1 chan interface{}
	ch1 = make(chan interface{})
	var b interface{}
	go func() {
		b = <-ch1
	}()
	ch1 <- stu1
	fmt.Printf("%T,%v", b, b)
}
