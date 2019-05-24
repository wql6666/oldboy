package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//计算两个大数的和，这两个大数会超过int64的范围
func main() {
	fmt.Println("请输入两个大数，中间以+连接：")
	line, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		fmt.Println("read from console err", err)
	}
	str := string(line)
	strNums := strings.Split(str, "+")
	if len(strNums) != 2 {
		fmt.Println("输入有误，请按格式输入：")
		return
	}

	strNum1 := strings.TrimSpace(strNums[0])
	strNum2 := strings.TrimSpace(strNums[1])
	count(strNum1, strNum2)

}
func count(str1, str2 string) {
	//判断长度，并在短的前边补”0“
	length1 := len(str1)
	length2 := len(str2)
	if length1 > length2 {
		for i := 0; i < length1-length2; i++ {
			str2 = "0" + str2

		}
	} else if length1 < length2 {
		for i := 0; i < length2-length1; i++ {
			str1 = "0" + str1
		}
	}
	var str1Result, str2Result []int32
	for _, v := range str1 {
		str1Result = append(str1Result, v-'0')
	}
	for _, v := range str2 {
		str2Result = append(str2Result, v-'0')
	}
	var left int32
	var resultArray []int32

	for i := 0; i < len(str1); i++ {
		result := str1Result[len(str1Result)-1-i] + str2Result[len(str2Result)-1-i] + left
		if result >= 10 {
			left = 1
		} else {
			left = 0
		}
		result = result % 10
		resultArray = append(resultArray, result)
		if i == len(str1)-1 && left == 1 {
			resultArray = append(resultArray, 1)
		}
	}
	for _, v := range resultArray {
		defer fmt.Print(v)

	}
}
