package main

import (
	"fmt"
	"math"
)
//100-200内素数
func main() {
	for i := 100; i <= 200; i++ {
		if isprime(i) {
			fmt.Println(i)
		}
		/*	for j := 2; j <= i; j++ {
			if j == i {
				fmt.Println(i)
			}
			if i%j == 0 && j < i {
				break
			}
		}*/
	}

}

func isprime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true

}
