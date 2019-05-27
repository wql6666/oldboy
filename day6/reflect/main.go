package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var a = 100
	//type1 := reflect.TypeOf(a)
	//addr1 := reflect.ValueOf(a).Interface()
	//kind := reflect.ValueOf(a).Kind()
	//fmt.Println(type1, addr1, kind)
	//fmt.Printf("%T", addr1)
	a := 10
	test(&a)
}
func test(a *int) {
	val := reflect.ValueOf(a)
	val.Elem().SetInt(100)
	fmt.Println(*a)
}




