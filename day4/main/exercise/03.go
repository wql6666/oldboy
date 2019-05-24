package main

import "fmt"
//判断是不是回文字符串
func main() {
	fmt.Println("qingshuruyigezifu")
	str:=""
	falg:=true
	fmt.Scan(&str)
	//注意string，[]byte,[]rune的区别
	t:=[]rune(str)
	fmt.Printf("%T，%c",t,t)
	//用len(str)/2减少循环次数。
	for i := 0; i <len(t)/2; i++ {
		if t[i]!=t[len(t)-1-i]{
			falg=false
		}
	}
	if falg==true{
		fmt.Println("shide",str)
	}else {
		fmt.Println("bushide")
	}
	//fmt.Println("------")
	//fmt.Println('中','国','人','民')
	//fmt.Println([]byte('中'))
}
