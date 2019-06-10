package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	test1()
	test()
}

func test() {
	s := os.Args
	fmt.Println("长度", len(s))
	for i, v := range s {
		fmt.Printf("Arg[%d]=%s\n", i, v)
	}
}

func test1() {
	var stringCommand string
	var intCommand int
	boolCommand:=flag.Bool( "b", false, "bool shiyong")
	flag.StringVar(&stringCommand, "s", "hei", "string shiyong")
	flag.IntVar(&intCommand, "i", 10, "int shiyong")
	flag.Parse()
	fmt.Println(*boolCommand)
	fmt.Println(stringCommand)
	fmt.Println(intCommand)
}
