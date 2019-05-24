package main

import "fmt"

func main() {
	slice := make([]map[int]int, 5)
	//if slice[0]==nil{
	//	slice[0]=make(map[int]int)
	//}
	for i := 0; i < 5; i++ {
		slice[i] = make(map[int]int)
	}
	for _, v := range slice {
		if v == nil {
			v = make(map[int]int)
		}
	}
	//fmt.Printf("%T",v)

	slice[0][0] = 100
	slice[1][1] = 200
	fmt.Println(slice)
	var a []int
	var b map[int]int
	var c chan int
	var d *int
	fmt.Println(a == nil,b==nil,c==nil,d==nil)
}
