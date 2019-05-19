package main

import "fmt"

func add(a, b *int) {
	*a,*b=*b,*a
	fmt.Println(*a,*b)
}
func main() {
	c:=add
fmt.Println(c,add)
a:=1000
b:=2000
c(&a,&b)
fmt.Println(a,b)



}
