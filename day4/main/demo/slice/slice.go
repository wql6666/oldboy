package main

import "fmt"

func main() {
	s1:=make([]int,10)
	var arr1=[5]int{}
	fmt.Println(s1)
	fmt.Println(arr1)
	modify(s1)
	modify2(arr1)
	fmt.Println(s1)
	fmt.Println(arr1)
}

func modify(s1 []int) {
	s1[0]=100
}

func modify2(arr [5]int) {
	arr[0]=100
}