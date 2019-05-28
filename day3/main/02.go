package main

import (
	"fmt"
	"strconv"
)
//水仙花数
func main() {
	for i := 100; i <= 999; i++ {
		var result = 0

		for _, v := range strconv.Itoa(i) {

			//fmt.Println(v)
			result += int(v-'0') * int(v-'0') * int(v-'0')
		}
		if result == i {
			fmt.Println(i)
		}
	}
	//fmt.Println('0')
}
