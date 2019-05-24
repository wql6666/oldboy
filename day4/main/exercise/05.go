package main

import "fmt"

func main() {
	for i := 2; i <= 1000; i++ {
		if i == sum(primeFactor(i)) {
			fmt.Println(i)
		}

	}

}

func primeFactor(n int) []int {
	var result []int
	flag := false
	for i := 2; i < n; i++ {
		for n != i {
			if n%i == 0 {
				result = append(result, i)
				n = n / i
				flag = true
			} else {
				break
			}
		}

	}
	if flag {
		result = append(result, n)
	}
	return result
}

func sum(array []int) int {
	var sumArray int
	for _, v := range array {
		sumArray += v
	}
	return sumArray

}
