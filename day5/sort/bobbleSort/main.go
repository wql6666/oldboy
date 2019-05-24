package main

import (
	"fmt"
)
func main() {
	var arr = [...]int{2, 432,999, 5, 1008,7,424, 56,1,31, 44, 4, 32, 76, 545, 99}
	sort1(arr[:])
	fmt.Println()

}

//冒泡排序：相邻两个比较
func sort1(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j]<a[j-1]{
				a[j],a[j-1]=a[j-1],a[j]
			}
		}
	}
	fmt.Println(a)
}