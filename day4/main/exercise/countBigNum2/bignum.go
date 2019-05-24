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
	result := count(strNum1, strNum2)
	fmt.Printf("%v", result)
	fmt.Println()

}
func count(str1, str2 string) string {
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
	var left uint8
	var result string
	index := (len(str1) - 1)
	for index >= 0 {
		c3 := str1[index] - '0' + str2[index] - '0' + left
		if c3 >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 = c3%10 + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index--
	}
	if left == 1 {
		result = fmt.Sprintf("1%s", result)
	}
	return result
}
