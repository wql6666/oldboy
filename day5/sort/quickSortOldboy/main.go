package main

import (
	"fmt"
)

func main() {
	var arr = [...]int{2, 432, 999, 5, 3, 89, 66, 1008, 7, 424, 56, 1, 31, 44, 4, 32, 76, 545, 99}
	//sort1(arr[:], 0, len(arr))
	fmt.Println(arr)
	sort2(arr[:], 0, len(arr)-1)
	fmt.Println(arr)
}

//快速排序：


func sort2(a []int, left, right int) {
	if left >= right {
		return
	}
	val:=a[left]
	k := left
	for i := left+1; i <= right; i++ {
		if a[i] < val {
			a[k] = a[i]
			a[i] = a[k+1]
			k++
		}

	}
	a[k] = val
	sort2(a, left, k-1)
	sort2(a, k+1, right)

}
