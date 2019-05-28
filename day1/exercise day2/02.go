package main

import "fmt"
//水仙花数
func main() {
	for i := 100; i <= 999; i++ {
		a:=i%10
		b:=i/100
		c:=(i/10)%10
		if (a*a*a+b*b*b+c*c*c)==i{
			fmt.Println(i)
		}

	}
	m,n:=0,0
	fmt.Scanf("%d%d",&m,&n)
fmt.Println(m,n)
}

func a(m ,n int)  {

}