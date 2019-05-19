package main

import "fmt"

type add_func func(int ,int)int

func sub(a, b int) int {
	return a-b
}

func operator(op func(a,b int)int, a,b int) int {
	return op(a,b)

}


func main() {
	c:=sub
	fmt.Println(c)
	result:=operator(c,100,200)
	fmt.Println(result)
}
