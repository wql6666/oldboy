package main

import (
	"fmt"
	"strings"
)

func main() {
	fun1:=makeSuffixFunc(".txt")
	fun2:=makeSuffixFunc(".jpg")
	fmt.Println(fun1("hello"))
	fmt.Println(fun2("qq"))
}

func makeSuffixFunc(suffix string)func(string)string{
	return func(name string) string {
		if !strings.HasSuffix(name,suffix){
			return name+suffix
		}else{
			return name
		}
	}

}