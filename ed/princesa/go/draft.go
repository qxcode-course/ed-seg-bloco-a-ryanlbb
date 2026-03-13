package main

import (
	"fmt"
)

func main() {
	var N, espada int
	fmt.Scan(&N, &espada)

	fila := make([]int, N)

	//INDICE de quem tem a espada e de quem vai morrer
	espada -= 1
	vitima := (espada + 1) % len(fila)

	for i := range fila {
		fila[i] = i + 1
	}

	vivos := len(fila)

	for vivos > 1 {
		printarFila(fila, espada)
		fila[vitima] = 0
		vivos--
		espada = proxVivo(fila, espada)
		vitima = proxVivo(fila, espada)
		
	}

	printarFila(fila, espada)
}

func proxVivo(fila []int, espadaPosic int) (vivo int){
	vivo = (espadaPosic + 1) % len(fila)
	for fila[vivo] == 0 {
		vivo = (vivo + 1) % len(fila)
	}

	return vivo
}

func printarFila(fila []int, espada int) {
	fmt.Print("[ ")
	for i, v := range fila {
		if v != 0 {
			if espada == i {
				fmt.Printf("%d> ", v)
			} else {
				fmt.Printf("%d ", v)
			}
		}
	}
	fmt.Println("]")
}
