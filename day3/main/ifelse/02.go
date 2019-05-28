package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("A")
		}
		fmt.Println()
	}
	test3()
}

func test3() {
	i := 0
lable:
	for ; i <= 10; i++ {
		fmt.Println("lalal")
		i++
		goto lable
		fmt.Println("hehe")

	}

}
