package main

import "fmt"

//find the longest uniq sub string
//every ch have is own start
func main() {
	s := "helloworlddjfkajkjkkjjkj"
	fmt.Println(findStr(s))
}

func findStr(s string) int {
	map1 := make(map[rune]int)
	maxLen := 0
	start := 0
	for i, ch := range s {
		lastI, ok := map1[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		map1[ch] = i
	}
	fmt.Println(len(s))
	fmt.Println(start, maxLen)
	//	return s[start : start+maxLen]
	return maxLen
}
