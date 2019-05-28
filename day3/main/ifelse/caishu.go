package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		fmt.Println("请输入一个整数：")
		n := 0
		fmt.Scanf("%d", &n)
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(101)
		fmt.Println("num为:", num)
		switch {
		case n == num:
			fmt.Println("猜对了")
		case n > num:
			fmt.Println("大了")
		default:
			fmt.Println("小了")

		}
	}
}
