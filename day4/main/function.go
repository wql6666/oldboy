package main

import "fmt"

func add(a, b int)int {
	return a+b
}
//

func main() {
	c:=add
	fmt.Printf("%T\n",c)
	sum:=c(10,20)
	fmt.Println(sum)

}
