package main

import "fmt"

func main() {
	i:=1
	defer add(&i)
	defer add1(i)
	i++
	fmt.Println(i)
	
}

func add(a *int)int {
	 *a=*a+1
	 fmt.Println(*a)
	 return *a

}
func add1(a int) int {
	a=a+1
	fmt.Println(a)
	return a
}