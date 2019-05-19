package main

import "fmt"

func main() {
	a,b:=10,20
	result:= func(a,b int)int {
		return a+b
	}(a,b)
	fmt.Println(result)
}
