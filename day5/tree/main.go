package main

import "fmt"

type student struct {
	name  string
	age   int
	score float32
	left  *student
	right *student
}

func main() {
	var root *student = &student{
		name:  "root",
		age:   20,
		score: 90.1,
	}

	var left1 *student = &student{
		name:  "left1",
		age:   20,
		score: 90.1,
	}

	var right1 *student = &student{
		name:  "right1",
		age:   20,
		score: 90.1,
	}

	var left2 *student = &student{
		name:  "left2",
		age:   20,
		score: 90.1,
	}
	var right2 *student = &student{
		name:  "right2",
		age:   20,
		score: 90.1,
	}
	root.left = left1
	root.right = right1
	left1.left = left2
	left1.right = right2
	trans(root)
}

func trans(p *student) {
	for p == nil {
		return
	}
	fmt.Println(p)
	trans(p.left)
	trans(p.right)

}
