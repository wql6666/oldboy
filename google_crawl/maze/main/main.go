package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

func main() {
	var fileNmae = "google_crawl/maze/main/mazi.in"
	maze := readMaze(fileNmae)
	start := point{0, 0}
	end := point{len(maze) - 1, len(maze[0]) - 1}
	steps := walk(maze, start, end)

	for _, v := range steps {
		for _, v1 := range v {
			fmt.Printf("%3d ", v1)
		}
		fmt.Println()
	}
}

func readMaze(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var row, col int
	fmt.Fscanf(file, "%d%d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}
func (p point) at(m [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(m) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(m[p.i]) {
		return 0, false
	}
	return m[p.i][p.j], true

}

func (p point) add(r point) point {
	var next = point{p.i + r.i, p.j + r.j}
	return next
}

//var step [4]point
var dirs = [4]point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	q := []point{start}
	for len(q) > 0 {
		cur := q[0] //walk at queue's first element
		q = q[1:]
		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)
			val, ok := next.at(maze)

			if val == 1 || !ok {
				continue
			}
			if next == start {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			curVal, _ := cur.at(steps)
			steps[next.i][next.j] = curVal + 1
			q = append(q, next)

		}

	}
	return steps

}
