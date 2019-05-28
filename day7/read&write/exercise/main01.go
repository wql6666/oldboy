package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//test()
	//test1()
	//test2()
	test3()//写入文件
}

func test() {

	file, err := os.Open("./day7/read&write/exercise/exercise.txt")
	if err != nil {
		fmt.Println("read file failed err=", err)
		return
	}
	defer file.Close()
	var (
		letterCount int
		numCount    int
		spaceCount  int
		otherCount  int
	)
	reader := bufio.NewReader(file)
	for ; ; {

		str, err := reader.ReadString('\n')
		fmt.Println(str)

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("radeline err=", err)
			return
		}

		for _, v := range []rune(str) {
			switch {
			case (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z'):
				letterCount++

			case v <= '9' && v >= '0':
				numCount++
			case v == ' ' || v == '\t':
				spaceCount++
			default:
				otherCount++
			}
		}
		fmt.Println("------")
	}

	fmt.Printf("lettercount=%d\tnumCount=%d\t spaceCount=%d\totherCount=%d",
		letterCount, numCount, spaceCount, otherCount)
}

func test1() {
	content, err := ioutil.ReadFile("./day7/read&write/exercise/exercise.txt")
	if err != nil {
		fmt.Println("ioutil read file failed err=", err)
		return
	}
	fmt.Println("content is:", string(content))
}

func test2() {
	file, err := os.Open("/root/桌面/go1.12.5.linux-amd64.gz")
	if err != nil {
		fmt.Println("open file failed err=", err)
	}
	gzreader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println("gizp file failed err=", err)
	}
	reader := bufio.NewReader(gzreader)
	str := make([]byte, 1000)
	for ; ; {
		str, err = reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read byte err=", err)
		}
		str = append(str, str...)
		//fmt.Println(str)

	}
	fmt.Println(str)
	err = ioutil.WriteFile("/root/桌面/gz.txt", str, 0777)
	if err != nil {
		fmt.Println("write file failed err=", err)
	}

}

func test3() {
	file, err := os.OpenFile("/root/桌面/a.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("hello world,i am coming")
	writer.Flush()

}
