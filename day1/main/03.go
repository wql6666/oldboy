package main

import "fmt"

func main() {
	var a =8
	var b =make(chan int)
	fmt.Println(a,b)
	modify(&a)
	fmt.Println(a)
	modify1(a)
	fmt.Println(a)
}

func modify(a *int) {
	*a=10
	return
}
func modify1(a int) {
	a=100
	return
}