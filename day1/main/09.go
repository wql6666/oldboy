package main

import "fmt"

func main() {
	a:="haha\n/n"
	b:=`aaa\n\n`
	c:='a'
	d:='o'
	fmt.Println(1,a,b,c)
	fmt.Println(2,a+b,b+b,b+a)
	fmt.Println(d)
}
