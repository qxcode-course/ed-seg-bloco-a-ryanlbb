package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

// encontra vizinho valido, poe no map de visitados e de casas anteriores
func getValidNeig(grid [][]rune, currentPos Pos, visited map[Pos]bool, previousPos map[Pos]Pos) []Pos {
	valid := []Pos{}

	for _, neig := range currentPos.getNeig() {
		if (match(grid, neig, ' ') || match(grid, neig, 'F')) && !visited[neig] {
			visited[neig] = true
			previousPos[neig] = currentPos
			valid = append(valid, neig)
		}
	}

	return valid
}

func (p Pos) getNeig() []Pos {
	return []Pos{
		{p.l + 1, p.c},
		{p.l - 1, p.c},
		{p.l, p.c + 1},
		{p.l, p.c - 1},
	}
}

func inside(grid [][]rune, pos Pos) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

func match(grid [][]rune, pos Pos, char rune) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == char
}

func search(grid [][]rune, startPos, endPos Pos) {
	previusPos := make(map[Pos]Pos)
	visited := make(map[Pos]bool)
	visited[startPos] = true
	queue := []Pos{startPos}
	idx := 0

	for idx < len(queue) {
		curr := queue[idx]
		idx++
		if curr == endPos {
			break
		}
		validNeigs := getValidNeig(grid, curr, visited, previusPos)
		queue = append(queue, validNeigs...)
	}

	voltar(grid, startPos, endPos, previusPos)
}

func voltar(grid [][]rune, startPos, endPos Pos, previousPos map[Pos]Pos) {
	for pos := endPos; pos != startPos; pos = previousPos[pos] {
		grid[pos.l][pos.c] = '.'
	}
	grid[startPos.l][startPos.c] = '.'
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nl, nc int
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	mat := make([][]rune, nl) // Inicializa a matriz de runes

	// Carregando matriz
	for i := range nl {
		scanner.Scan()
		line := scanner.Text()
		mat[i] = []rune(line)
	}

	var inicio, fim Pos

	// Procurando inicio e fim e colocando ' ' nas posições iniciais
	for l := range nl {
		for c := range nc {
			if mat[l][c] == 'I' {
				mat[l][c] = ' '
				inicio = Pos{l, c}
			}
			if mat[l][c] == 'F' {
				mat[l][c] = ' '
				fim = Pos{l, c}
			}
		}
	}

	search(mat, inicio, fim)

	for _, line := range mat {
		fmt.Println(string(line)) // Converte o slice de runes de volta para string para imprimir
	}
}
