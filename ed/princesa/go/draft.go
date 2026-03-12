package main

import (
	"fmt"
)

func main() {
	var N, espada int
	fmt.Scan(&N, &espada)

	vaiMorre := espada
	fila := make([]int, N)

	for i := range fila {
		fila[i] = i + 1
	}

	mortos := 0

	for mortos < N - 2 {
		fmt.Print("[ ")

        // primeira iteracao, imprime todos
		if mortos == 0 {
			for i, v := range fila {
				fmt.Print(v)

				if espada == i+1 {
					fmt.Print("> ")
				}
			}
            fmt.Println(" ]")
            fila[vaiMorre] = 0
            mortos++
		}

		for i, v := range fila {
			if v != 0 {
				fmt.Print(v)
			}

			if espada == i+1 {
				fmt.Print("> ")
			}

			fila[vaiMorre] = 0
		}
		mortos++
		fmt.Println(" ]")
	}
}
