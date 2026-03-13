package main

import (
	"fmt"
	"slices"
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

	mortos := 0

    //mostra todos
    fmt.Print("[ ")
	for i, v := range fila {
	    if espada == i{
			fmt.Printf("%d> ", v)
		} else {
            fmt.Printf("%d ", v)
        }
	}
    fmt.Println("]")
    fila[vitima] = 0
    mortos++
    
    //matança
	for mortos < len(fila) {
		
	}
}

func matar (espadaPosic int, vet []int) {
	vitima := (espadaPosic + 1) % len[vet]
	
	for {
			
	}
}
