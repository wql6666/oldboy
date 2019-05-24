package main

import "fmt"

func main() {
	map1 := make(map[int]string)
	map2 := make(map[string]int)
	map1[1] = "hello"
	map1[2] = "world"
	map1[3] = "hei"
	map1[4] = "man"
	for k, v := range map1 {
		map2[v] = k
	}
	fmt.Println(map2)
}
