package main

import "fmt"

func main() {
	arr()
}

func arr() {
/*	var a [10]int
	a[0] = 1
	b := 9
	a[b] = 2
	fmt.Println(a)
	*/
	for i := 1; i <= 50 ;i++{
		fmt.Print(fb(i),"\t")
	}

}

func fb(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	var a, b = 1, 1
	var sum int
	if n >= 3 {
		for i := 3; i <= n; i++ {
			sum = a + b
			a = b
			b = sum
		}

	}
	return sum

}
