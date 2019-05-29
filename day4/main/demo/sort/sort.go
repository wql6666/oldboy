package main

import (
	"fmt"
	"sort"
)

func main() {
	//testexaple()
	//testString()
	//testFloat()
	testSearchInts()
}

func test() {
	var a = [...]int{432, 342, 3, 231, 55, 2}
	sort.Ints(a[:])
	fmt.Println(a)
}

func testString() {
	var a =[...]string{"a","lala","A","hello"}
	sort.Strings(a[:])
	fmt.Println(a)
}
func testFloat() {
	var a =[...]float64{1.2,3232.44,3.2,0.7}
	sort.Float64s(a[:])
	fmt.Println(a)

}

func testSearchInts() {
	var a =[]int{22,234,42,2,767,34,4}
	sort.Ints(a)
	fmt.Println(a)
	index:=sort.SearchInts(a,34)
	fmt.Println(index)
}