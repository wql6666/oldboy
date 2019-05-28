package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n string
	fmt.Scanf("%v", &n)
	result,err:=strconv.Atoi(n)
	if err!=nil{
		fmt.Println("can not convert to int",err)
	}else {
		fmt.Printf("%v\n", result)
		fmt.Printf("%T", result)
	}
}
