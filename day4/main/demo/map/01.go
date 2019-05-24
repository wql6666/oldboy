package main

import "fmt"

func main() {
	var a []int
	// a=make([]int,5)
	a=[]int{0,0,0,0,0}
	a[0]=1
	fmt.Println(a)



}
func test() {
	var a *int
	a=new(int)
	*a=10
	fmt.Println(*a)

	var b *[]int
	//b=make([]int,8)
	b=new([]int)
	*b=[]int{1,2,3}
	(*b)[0]=1
	fmt.Println(b)

}
