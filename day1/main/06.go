package main

import "fmt"

var b string

func main() {
	b = "H"
	fmt.Println(b)
	f1()
	test()
}

func f1() {
	b := "K"
	fmt.Println(b)
	f2()
}

func f2() {
	fmt.Println(b)
}

func test() {
	var a int8 = 10
	var b int32 = 100
	b = int32(a)
	fmt.Println(b)
	var c int=20
	var d int32=30
	fmt.Println(c+int(d))
}
