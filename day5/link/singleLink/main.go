package main

import (
	"fmt"
	"math/rand"
)

type student struct {
	name  string
	age   int
	score float32
	next  *student
}

func main() {
	var head *student = &student{
		name:  "head",
		age:   20,
		score: 90.1,
	}
	var p = &student{}
	p.next = head
	insertTail(head)
	trans(head)
	fmt.Println()
	//head = insertHead(head)
	//trans(head)
	fmt.Println("--------")
	//insertHead2(&head)
	//trans(head)

	linkDelete(head)
	trans(head)
	fmt.Println("++++++++")
	var stu6 = &student{
		name:  "stu6",
		age:   100,
		score: 100,
	}
	insert(head,stu6)
	trans(head)

}

func trans(p *student) {
	for p != nil {
		fmt.Println(p)
		p = p.next
	}

}

func insertTail(head *student) {

	var tail = head

	for i := 0; i <= 10; i++ {
		var stu *student = &student{
			name:  fmt.Sprintf("stu%d", i),
			age:   rand.Intn(10) + 15,
			score: rand.Float32() * 100,
		}
		tail.next = stu
		tail = stu
	}
}

func insertHead(head *student) *student {

	for i := 0; i <= 10; i++ {
		var stu *student = &student{
			name:  fmt.Sprintf("stu%d", i),
			age:   rand.Intn(10) + 15,
			score: rand.Float32() * 100,
		}
		stu.next = head
		head = stu
	}
	//trans(head)
	return head
}

func insertHead2(head **student) {

	for i := 0; i <= 10; i++ {
		var stu *student = &student{
			name:  fmt.Sprintf("stu%d", i),
			age:   rand.Intn(10) + 15,
			score: rand.Float32() * 100,
		}
		stu.next = *head
		*head = stu
	}
}
//TODO 删除不了头节点的。
func linkDelete(p *student) {
	var previous *student = p
	for p != nil {
		if (*p).name == "stu6" {
			previous.next = (*p).next
			break
		}
		previous = p
		p =p.next
	}

}

func insert(p ,newNode *student ) {

	for p != nil {
		if p.name == "stu5" {
			newNode.next = p.next
			p.next = newNode
			break
		}
		p = p.next

	}

}
