package main

import "fmt"


var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	fmt.Println(a)
}
func m() {
	a := "O"//这是声明了一个局部变量
	//a="O" 注意区别
	fmt.Println(a)
}

