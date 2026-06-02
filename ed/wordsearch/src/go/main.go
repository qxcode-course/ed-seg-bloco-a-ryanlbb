package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	for i := range grid {
		for j := range grid[i] {
			if word[0] == grid[i][j] {
				if search(grid, word, 0, i, j) {
					return true
				}
			}
		}
	}
	return false
}

func search(grid [][]byte, word string, strIdx, i, j int) bool {
	if i >= len(grid) || j >= len(grid[0]) || i < 0 || j < 0 || grid[i][j] != word[strIdx] {
		return false
	}
	if strIdx == len(word)-1 {
		return true
	}

	/* marca a posic que esta atualmente
	como visitada pra nao checar aqui novamente*/
	anteriorByte := grid[i][j]
	grid[i][j] = '#'

	solve := search(grid, word, strIdx+1, i-1, j) ||
		search(grid, word, strIdx+1, i+1, j) ||
		search(grid, word, strIdx+1, i, j-1) ||
		search(grid, word, strIdx+1, i, j+1)

	// volta o estado anteriror
	grid[i][j] = anteriorByte

	return solve
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
