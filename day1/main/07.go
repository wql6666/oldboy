package main

import "fmt"

func main() {
	var a int32=100
	var b int16
	b=a
	b=int16(a)
	fmt.Println(a,b)
}
