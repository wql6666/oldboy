package main

import "fmt"

//阶乘之和
func main() {
	n := 0
	sum := 0
	s := 1
	fmt.Println("请输入一个整数n:")
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		s = i * s
		sum += s
	}
	fmt.Println(sum)

}
