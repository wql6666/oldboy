package main

import "fmt"
//区分因数和质因数
func main() {
	for i := 1; i <= 1000; i++ {
		if i == factor(i) {
			fmt.Println(i)
		}
	}
}

func factor(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}

	}
	return sum

}
