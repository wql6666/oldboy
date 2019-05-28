package main

import (
	"fmt"
)

func main() {
	reverseInt(5423524352)
	fmt.Println(reverseInt1(5423524352))
}
func reverseInt(z interface{}) {
	var result string
	str := fmt.Sprintf("%v", z)
	strLen := len(str)
	for i := 1; i <= strLen; i++ {
		result += string(str[strLen-i])
	}
	fmt.Println(result)
}

func reverseInt1(x int) int {
	var num = 0
	var newNum = 0
	for x != 0 {
		a := x % 10
		newNum = num*10 + a
		num = newNum
		x = x / 10

	}
	return num
}
