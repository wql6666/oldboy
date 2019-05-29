package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num - 1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		var m map[string]int
		m["alan"] = 10
		fmt.Println(m)

	}()
	go cal()
	//for ; ;  {
	//	fmt.Println("--")
	//}
	fmt.Println("hahah")
	time.Sleep(time.Second * 100)
}
func cal() {
	for ; ; {
		i := 1
		fmt.Println(i)
		i++
		time.Sleep(time.Second)
	}

}
