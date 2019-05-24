package main

import (
	"fmt"
)

func main() {
	var arr = [...]int{2, 432, 999, 5, 1008, 7, 424, 56, 1, 31, 44, 4, 32, 76, 545, 99}
	sort1(arr[:])
	fmt.Println()

}

//插入排序,前边是一个有序序列
//从后边无序数列中依次选出一个数
//这个数和依次和前边的有序数列进行比较，大于则跳出循环，若是小于则交换位置。
func sort1(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] > a[j-1] {
				break
			}else {
				a[j],a[j-1]=a[j-1],a[j]
			}
		}
	}
	fmt.Println(a)
}
