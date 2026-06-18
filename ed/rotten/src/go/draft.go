package main

import (
	"fmt"
)

type Pos struct {
	x int
	y int
}

func isOrange(grid [][]int, p Pos) bool {
	if p.x < 0 || p.y < 0 || p.x >= len(grid) || p.y >= len(grid[0]) {
		return false
	}
	if grid[p.x][p.y] != 1 {
		return false
	}

	return true
}

func rotten(grid [][]int) int {
	time := 0
	oranges := 0
	queue := []Pos{}
	for l := range grid {
		for c := range grid[l] {
			if grid[l][c] == 2 {
				queue = append(queue, Pos{l, c})
			}
			if grid[l][c] == 1 {
				oranges++
			}
		}
	}

	batch := len(queue)
	for len(queue) > 0 {

		oldPos := queue[0]
		queue = queue[1:]

		up := Pos{oldPos.x, oldPos.y + 1}
		down := Pos{oldPos.x, oldPos.y - 1}
		right := Pos{oldPos.x + 1, oldPos.y}
		left := Pos{oldPos.x - 1, oldPos.y}

		if isOrange(grid, up) {
			grid[up.x][up.y] = 2
			queue = append(queue, up)
			oranges--
		}
		if isOrange(grid, down) {
			grid[down.x][down.y] = 2
			queue = append(queue, down)
			oranges--
		}
		if isOrange(grid, right) {
			grid[right.x][right.y] = 2
			queue = append(queue, right)
			oranges--
		}
		if isOrange(grid, left) {
			grid[left.x][left.y] = 2
			queue = append(queue, left)
			oranges--
		}

		batch--
		if batch == 0 && len(queue) > 0 {
			time++
			batch = len(queue)
		}
	}

	if oranges > 0 {
		return -1
	}
	return time
}

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	oranges := make([][]int, m)

	for l := range oranges {
		oranges[l] = make([]int, n)
	}

	for l := range oranges {
		for c := range oranges[l] {
			fmt.Scan(&oranges[l][c])
		}
	}

	fmt.Println(rotten(oranges))
}
