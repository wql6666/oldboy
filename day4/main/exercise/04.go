package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("请输入一个字符串:")
	//result := ""
	//fmt.Scan(&result)

	//reader := bufio.NewReader(os.Stdin)
	result,_, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		fmt.Println("read from console err", err)
		return
	}
	lc, nc, sc, oc := count(string(result))
	fmt.Printf("字母次数为%d\n数字次数为%d\n空格次数为%d\n其他次数为%d\n", lc, nc, sc, oc)

}

func count(str string) (letterCount, numCount, spaceCount, otherCount int) {
	t := []rune(str)
	for _, v := range t {
		switch {
		case (v >= 'a' && v <= 'z' )|| (v >= 'A' && v <= 'Z'):
			letterCount++
		case v >= '0' && v <='9':
			numCount++
		case v == ' ':
			spaceCount++
		default:
			otherCount++

		}
	}
	return

}
