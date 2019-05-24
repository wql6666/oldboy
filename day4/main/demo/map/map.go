package main

import "fmt"

func main() {
	//var a map[string]string
	a:=make(map[string]string,1)

	a["hello"]="world"
	length:=len(a)
	fmt.Println(length)
	fmt.Printf("%p",a)
	a["yo"]="girl"
	a["good"]="morning"
	a["hi"]="Bob"
	a["oh"]="no"
	a["hei"]="man"
	fmt.Println()
	fmt.Printf("%p",a)
	fmt.Println(a)
	var map2=map[string]string{
		"hello":"world",
	}
fmt.Println(map2)
}
