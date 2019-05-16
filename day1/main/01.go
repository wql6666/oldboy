package main

import (
	"fmt"
	a "oldboy/day1/add"
	"oldboy/day1/test"
)

func main() {
	//n := 0
	//fmt.Println("请输入一个整数：")
	//fmt.Scan(&n)
	//list(n)
	fmt.Println(a.Name,a.Age)
	fmt.Println(test.B)
}

func list(n int) {

	for i := 0; i <= n; i++ {
		fmt.Printf("%d+%d=%d\n", i, n-i, n)

	}
}
