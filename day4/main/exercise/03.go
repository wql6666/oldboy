package main

import "fmt"

func main() {
	fmt.Println("qingshuruyigezifu")
	str:=""
	falg:=true
	fmt.Scan(&str)
	for i := 0; i <= len(str)-1; i++ {
		if str[i]!=str[len(str)-1]{
			falg=false
		}
	}
	if falg==true{
		fmt.Println("shide",str)
	}
}
