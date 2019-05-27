package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	stu1 := student{}
	judge(2, "haha", 2.2, true, 1, false, 0, stu1)
}

func judge(args ...interface{}) {
	for i, v := range args {
		switch v.(type) {
		case int:
			fmt.Printf("第%d个参数:%v的type is int\n", i, v)
		case string:
			fmt.Printf("第%d个参数:%v的type is string\n", i, v)
		case bool:
			fmt.Printf("第%d个参数:%v的type is bool\n", i, v)
		case float64:
			fmt.Printf("第%d个参数:%v的type is float64\n", i, v)
		case student:
			fmt.Printf("第%d个参数:%v的type is student\n", i, v)

		default:
			fmt.Printf("第%d个参数:%v的type is unknow\n", i, v)
		}
	}
}
