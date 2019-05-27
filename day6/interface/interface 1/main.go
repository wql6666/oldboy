package main

import "fmt"

type writer interface {
	write()
}
type reader interface {
	read()
}

type readerWriter interface {
	writer
	reader
}

type f struct {
}

func (file *f) read() {
	fmt.Println("read book")

}
func (file *f) write() {
	fmt.Println("write book")

}

func Test(rw readerWriter) {
	fmt.Println("hahaha")

}

func main() {
	var file = f{}
	Test(&file)
	var a interface{}
	a = &file
	a, ok := a.(readerWriter)
	fmt.Println(a, ok)

}
