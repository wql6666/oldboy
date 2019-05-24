package main

import (
	"fmt"
)
func main() {
	var arr = [...]int{2, 432,999, 5, 1008,7,424, 56,1,31, 44, 4, 32, 76, 545, 99}
	sort1(arr[:])
	fmt.Println()
	sort2(arr[:])
	fmt.Println()
}

//属于选择排序，选出一个最小的或最大的
func sort1(arr []int)  {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)
}

func sort2(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[i]<arr[j]{
				arr[i],arr[j]=arr[j],arr[i]
			}
		}
	}
	fmt.Print(arr)
}