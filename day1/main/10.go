package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		a int=1000
		b string = "hello"
		c bool
		d byte
	)
	fmt.Printf("%v,%T,%T,%v\n", a, b, c, d)
	f, err := strconv.Atoi("111")
	fmt.Println(f, err)
	str:=fmt.Sprintf("%d",a)
	fmt.Println(str)
	fmt.Printf("%T",str)
}
