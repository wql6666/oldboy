package main

import (
	"fmt"
)

func main() {
	var arr = [...]int{2, 432, 999, 5, 3, 89, 66, 1008, 7, 424, 56, 1, 31, 44, 4, 32, 76, 545, 99}
	sort1(arr[:], 0, len(arr))
	fmt.Println(arr)
	sort2(arr[:], 0, len(arr))
	fmt.Println(arr)
}

//快速排序：
func sort1(a []int, left, right int) {

	if left >= right {
		return
	}
	//k是变化的left，下标
	k := left
	for i := left + 1; i < right; i++ {
		if a[i] < a[k] {
			a[k+1], a[i] = a[i], a[k+1]
			a[k], a[k+1] = a[k+1], a[k]
			k++
		}
	}
	sort1(a, left, k)
	sort1(a, k+1, right)

}

func sort2(a []int, left, right int) {
	if left >= right {
		return
	}
	n := left
	for i := left+1; i < right; i++ {
		if a[i] < a[n] {
			a[n] = a[i]
			a[i] = a[n+1]
			n++
		}

	}
	a[n] = a[left]
	sort2(a, left, n)
	sort2(a, n+1, right)

}
