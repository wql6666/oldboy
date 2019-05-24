package main

import "fmt"

func main() {
	f:=test()
	fmt.Println(f(1))
	fmt.Println(f(100))
	fmt.Println(f(1000))

}

func test()func(int) int{
	var x int
	return func(d int) int {
		x+=d
		return x
	}

}
