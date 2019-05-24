package main

import "fmt"
//用递归求阶乘之和
func main() {
	fmt.Println(sum(5))
}

func recusive(n int) int {
	//定义出口，不然就是死循环
	if n==1{
		return 1
	}
	return  recusive(n-1)*n
}

func sum(n int)int{
	sum:=0
	for i := 1; i <= n; i++ {
		sum=sum+ recusive(i)
	}
	return sum
}