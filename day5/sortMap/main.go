package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := make(map[int]int)
	map1[2] = 100
	map1[1] = 200
	map1[3] = 322
	map1[78] = 434
	map1[43] = 43
	map1[66] = 89
	sortMap(map1)
}

func sortMap(map1 map[int]int) {
	var s1 []int
	str := ""
	for k, _ := range map1 {
		s1 = append(s1, k)
	}
	sort.Ints(s1)
	for _, v := range s1 {
		//fmt.Println(v,map1[v])
		str += fmt.Sprintf("%d:%d\t", v, map1[v])
	}
	fmt.Println(str)
}
