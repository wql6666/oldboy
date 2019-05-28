package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	b := 5
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err recover")
		}
	}()
	fmt.Println(a[b])

}
