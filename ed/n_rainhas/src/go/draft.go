package main

import "fmt"

func queens(board [][]rune) int {
	cont := 0

	return cont
}

func main() {
	var n int
	fmt.Scan(&n)
	board := make([][]rune, n)
	for i := range board {
		board[i] = make([]rune, n)
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] = 'F'
		}
	}
	fmt.Println(board)
	fmt.Println(queens(board))
}

