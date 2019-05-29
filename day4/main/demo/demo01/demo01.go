package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	//for ; ; {
	//	testexaple()
	//	time.Sleep(time.Second)
	//}
	err:=initConfig()
	if err!=nil{
		panic("init config failed")
	}
	j := new(int)
	fmt.Println(j)
	ch1 := make(chan int)
	fmt.Println(ch1)
	var arr []int
	arr = append(arr, 9, 8)
	fmt.Println(arr)

}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	a := 0
	b := 10
	result := b / a
	fmt.Println(result)
}

func initConfig() (err error) {
	return errors.New("innit config err")

}