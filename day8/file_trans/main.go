package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	fileRead  = "/root/桌面/zhixing.jpg"
	fileWrite = "/root/桌面/write.jpg"
	temp      = "/root/桌面/temp.txt"
	bs        = make([]byte, 1) //temp
	readByte  = make([]byte, 1024)
)

func main() {
	data := read()
	fmt.Println(string(data))
	//write()
}

func read() []byte {
	file, err := os.OpenFile(temp, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic("open temp failed")
	}
	var n int
	for ; ; {
		n, err = file.Read(bs)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("temp file read failed", err)
			panic("temp file read failed")
		}
		fmt.Println("read temp file size=", n)
	}

	offset, err := strconv.Atoi(string(bs[:n]))
	if err != nil {
		fmt.Println("convert bs to int failed", err)
	}
	fileSrc, err := os.OpenFile(fileRead, os.O_CREATE|os.O_RDWR, 0777)
	_, err = fileSrc.Seek(int64(offset), 0)
	if err != nil {
		fmt.Println("seek position failed", err)
	}
	reader := bufio.NewReader(fileSrc)
	for ; ; {
		readByte, err = reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("read read byte failed")
		}
		if err == io.EOF {
			break
		}

	}
	return readByte

}

func write(b []byte) {
	//writefile,err:=os.OpenFile(fileWrite,os.O_CREATE|os.O_RDWR,0777)
	//if err!=nil{
	//	fmt.Println("open write file failed")
	//	return
	//}
	//writefile.Seek()

}
