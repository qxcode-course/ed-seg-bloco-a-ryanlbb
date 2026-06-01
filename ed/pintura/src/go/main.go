package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Não modifique a assinatura da função floodFill
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	for i := range image {
		for j := range image[i] {
			if i == sr && j == sc {
				toPaint := image[i][j]
				paint(image, toPaint, color, i, j)
			}
		}
	}

	return image
}

func paint(image [][]int, toPaint, color, i, j int) {
	if i < 0 || j < 0 || i >= len(image) || j >= len(image[0]) || image[i][j] != toPaint || image[i][j] == color {
		return
	}

	image[i][j] = color

	paint(image, toPaint, color, i-1, j)
	paint(image, toPaint, color, i+1, j)
	paint(image, toPaint, color, i, j-1)
	paint(image, toPaint, color, i, j+1)
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Fields(line)
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	image := make([][]int, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		rowStr := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc; j++ {
			row[j], _ = strconv.Atoi(rowStr[j])
		}
		image[i] = row
	}

	// Lê sr, sc e color
	scanner.Scan()
	params := strings.Fields(scanner.Text())
	sr, _ := strconv.Atoi(params[0])
	sc, _ := strconv.Atoi(params[1])
	color, _ := strconv.Atoi(params[2])

	result := floodFill(image, sr, sc, color)

	// Imprime a matriz resultante
	for _, row := range result {
		for j, val := range row {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}
