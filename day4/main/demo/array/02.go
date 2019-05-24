package main

import "fmt"

func main() {
	fmt.Print(fab1(20))
	var a =[5]int{1,2,3}
	var b [5]int=[5]int{1,2,3,4}
	var c =[...]int{1,2,3,4,5}
	var d=[...]int{1:200,3:400}
	fmt.Println()
	fmt.Println(a,b,c,d)
	test()
}

func fab1(n int) []int {
	s1 := make([]int, n)
	//var s1 []int
	//s1=append(s1,n)
	s1[0] = 1
	s1[1] = 1
	for i := 2; i < n; i++ {
		s1[i] = s1[i-1] + s1[i-2]
	}
	return s1
}

func test() {
	var s =[...][3]int{{1,2,3},{2,3,4}}
	fmt.Println(s)
}