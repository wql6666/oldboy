package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	srcFile,err:=os.Open("/workspace/go/src/oldboy/day7/read&write/hello.txt")
	if err!=nil{
		fmt.Println("open err",err)
	}
	defer srcFile.Close()
	n,err:=io.Copy(os.Stdout,srcFile)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("\n",n)

}
