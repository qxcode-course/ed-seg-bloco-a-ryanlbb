package main

import (
	"fmt"
)

func main() {
	var nFila, nDeixouFila int
	fmt.Scan(&nFila)
	fila := make([]int, nFila)
	for i := range nFila {
		var valor int
		fmt.Scan(&valor)
		fila[i] = valor
	}

	fmt.Scan(&nDeixouFila)
	deixaramFila := make(map[int]bool)
	for range nDeixouFila {
		var valor int
		fmt.Scan(&valor)
		deixaramFila[valor] = true
	}

	for i, v := range fila {
		if deixaramFila[v] == true {
			continue
		} else {
			fmt.Print(fila[i], " ")
		}
	}

	fmt.Println()
}
