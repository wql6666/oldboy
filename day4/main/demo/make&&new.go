package main

import "fmt"

func main() {
	test()
}

func test() {
	s1:=new ([]int)
	s2:=make([]int,10)
	stru1:=new(struct{})
	//stru2:=make(struct{})
	map1:=new(map[int]string)
	map2:=make(map[int]string)
	ch1:=new(chan int)
	ch2:=make(chan int)
	fmt.Printf("s1=%v\t,s2=%v\t,stru1=%v\t,map1=%v\t" +
		",map2=%v\t,ch1=%v\t,ch2=%v\t",s1,s2,stru1,map1,map2,ch1,ch2)
	//make([5]int)
	*s1=make([]int,5)
	(*s1)[0]=100
	s2[0]=100
	*map1=make(map[int]string)
	(*map1)[1]="good"
	map2[1]="good"
	*ch1=make(chan int)


	fmt.Println(s1,s2,map1,map2,ch1)
}
