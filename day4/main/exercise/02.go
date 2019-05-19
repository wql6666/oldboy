package main

import "fmt"

func main() {

	for i:=1;i<=1000;i++{
		var v int
		for _,v=range(factor(i)){
			v+=v
		}
		if i==v{
			fmt.Println(i)
		}
	}
}

func factor (n int)[]int  {
	var result []int
	  m:=n
	for i:=1;i<=m;i++{
		if n%i==0{
			result=append(result,i)
		//	fmt.Println("****",n)
		//	fmt.Println("&&&",result)
			n=n/i
		}

	}
	return result

}