package main

import "fmt"

type a interface {
}

func main() {
	var b int
	var a1 a
	a1 = b
	y, ok := a1.(string)
	if ok {
		fmt.Println(y)
	} else {
		fmt.Println("convert failed")
	}
}
