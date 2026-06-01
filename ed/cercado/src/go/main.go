package main

import (
	"bufio"
	"fmt"
	"os"
)

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' && (i == len(board)-1 || j == len(board[i])-1 || i == 0 || j == 0) {
				save(board, i, j)
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'
			case 'S':
				board[i][j] = 'O'
			}
		}
	}
}

func save(board [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) || board[i][j] != 'O' {
		return
	}

	board[i][j] = 'S'

	save(board, i-1, j)
	save(board, i+1, j)
	save(board, i, j-1)
	save(board, i, j+1)
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
