package main

import (
	"fmt"
)

func main() {
	//test('a')
	var a = []int{}
	test1(&a)

}

func test(character byte) byte {
	fmt.Println('A' - 'a')
	return 'A'
}

func test1(A *[]int) {
	for i, _ := range *A {
		for j := 1; j < i; j++ {
			if (*A)[j-1] > (*A)[j] {
				(*A)[j-1], (*A)[j] = (*A)[j], (*A)[j-1]
			}
		}
	}
	fmt.Println(*A)

}
