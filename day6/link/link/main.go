package main

import "fmt"

type linkNode struct {
	data interface{}
	next *linkNode
}

type link struct {
	head *linkNode
	tail *linkNode
}


func main() {
	var link link
	for i := 1; i <= 10; i++ {
		link.insertTail(fmt.Sprintf("student%d",i))

	}
	link.trans()
	//fmt.Println(*link.head)

	fmt.Println(*link.tail)
}

func(link *link) insertHead(data interface{}) {
	var node = &linkNode{
		data: data,
		//nil,
	}
	if link.tail == nil && link.head == nil {
		link.tail = node
		link.head = node
		return
	}
	node.next = link.head
	 link.head=node
}

func (link *link)insertTail(data interface{}) {
	var node = &linkNode{
		data: data,
		//nil,
	}
	if link.tail == nil && link.head == nil {
		link.head = node
		link.tail = node
		return
	}
	link.tail.next = node
	link.tail=node
}

func (link *link)trans() {
	for link.head != nil {
		fmt.Println(link.head.data)
		link.head = link.head.next
	}
}
