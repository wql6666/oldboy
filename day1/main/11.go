package main

import "fmt"

func main() {
	var (
		a string = "hello"
		b        = "world"
	)
	str := fmt.Sprintf("%v%v", a, b)
	fmt.Println(str)
	str1 := str[:5]
	fmt.Println(str1)
	str2 := str[5:]
	fmt.Println(str2)

	result := reverse(str)
	fmt.Println(result)
	//fmt.Printf("%c",str[1])
	result1 := reverse1(result)
	fmt.Println(result1)
}

func reverse(str string) string {
	var result string
	strLen := len(str)
	for i := 1; i <= strLen; i++ {
		result += fmt.Sprintf("%c", str[strLen-i])

	}

	return result
}

func reverse1(str string) string {
	var result []byte
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		result = append(result, str[strLen-i-1])
	}
	return string(result)
}
