package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {

		fmt.Println(fab(i))
	}
}

func fab(n int) int {
	if n==1||n==2{
		return 1
	}
	return fab(n-1)+fab(n-2)
}
