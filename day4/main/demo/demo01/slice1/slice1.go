package main

import "fmt"

func main() {
	//test()
	var a =[]int{1,2,3,5,5}
	var b = []int{0,0,0,0,0,6,7,8,9,10}
	num:=copy(a,b)
	fmt.Println(num)
	fmt.Println(b)
	fmt.Println(a)
str:="中国 hello world"
fmt.Printf("%T\t%p",str,&str)
fmt.Println()
str1:=[]rune(str)
str1[0]='0'
fmt.Println(string(str1))
fmt.Printf("%p\n%p",&str,str1)

}

func test() {
	var arr =[10]int{1,2,3,4,5,6,7,8,9,10}
	s1:=arr[1:5]
	fmt.Printf("%p\n%p\n",&arr[1],s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	s3:=arr[2:4]
	s2:=arr[2:]
	s1=append(s1,s2...)
	fmt.Println(len(s1),cap(s1))
	s4:=append(s1,s3...)
	fmt.Println(s1,"\n",s4)
	fmt.Printf("%p\n%p",s1,s4)
}