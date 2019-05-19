package main

import "fmt"

func main() {
	sum:=add(1,2,3,4,5,100)
	fmt.Println(sum)
	str:=contact("hello","world","haha","niao")
	fmt.Println(str)
}

func add(a int ,arg ...int)( sum int) {
	sum =a
	for _,v:=range arg{
		sum+=v
	}
	return

}
func contact(a string, arg ...string) string {

	for _,v:=range arg{
     a+=v
	}
	return a

	//return fmt.Sprintf("%s,%s,%s",a,arg[0],arg[1])

}