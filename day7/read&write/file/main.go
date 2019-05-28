package main

import (
	"fmt"
	"os"
)

func main() {
	file,err:=os.OpenFile("./day7/log.txt",os.O_CREATE|os.O_WRONLY,777)
	defer file.Close()
	if err!=nil{
		fmt.Println("open file err=",err)
		return
	}
	fmt.Fprintln(file,"err lalala")
	fmt.Fprintln(os.Stdout,"lala")
}
