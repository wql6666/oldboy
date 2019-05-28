package main

import "fmt"

func main() {
	var a int
	a = 2
	switch a {
	case 0,1,2:
		fmt.Println("wei1")
		fallthrough
	case 3:
		fmt.Println("wei2")
	default:
		fmt.Println("wei3")
	}
	test()
	test1()
}

func test() {
	switch {
	case true:
		fmt.Println("hah ")
	case false:
		fmt.Println("lala")

	}
}
func test1()  {
	switch i := 0; {
	case i>0:
		fmt.Println("haha")
	default:
		fmt.Println("shibai")


	}

}