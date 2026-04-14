package main

import (
	"fmt"
)

func printaFaltantes(vet []int, mapa map[int]bool) {
	falta := false
	for i := range mapa {
		if mapa[i] == false {
			falta = true
		}
	}

	if falta == false {
		fmt.Println("N")
		return
	}

    vetFaltantes := make([]int, 0)
    for _, v := range vet {
		if mapa[v] == false {
			vetFaltantes = append(vetFaltantes, v)
            }
		}
    
	for i, v := range vetFaltantes {		
		fmt.Print(v)
		if i < len(vetFaltantes)-1 {
				fmt.Print(" ")
		}
	}
    fmt.Println()
}

func printaRepetidos(vet []int, mapa map[int]int) {
	repete := false
	for i := range mapa {
		if mapa[i] > 1 {
			repete = true
		}
	}

	if repete == false {
		fmt.Println("N")
		return
	}

	for i, v := range vet {
		if mapa[v] > 1 {
			fmt.Print(v)
			if i < len(vet)-2 {
				fmt.Print(" ")
			}
			mapa[v]--
		}
	}
	fmt.Println()
}

func main() {
	var nFigs, nBaruel int
	fmt.Scan(&nFigs)

	baruelFalta := make(map[int]bool)
	baruelRepetidos := make(map[int]int)
	figs := make([]int, nFigs)

	// todas as fig original
	for i := range figs {
		figs[i] = i + 1
		baruelFalta[i+1] = false
	}

	fmt.Scan(&nBaruel)
	baruel := make([]int, nBaruel)

	// preenche todos os vetores
	for i := range nBaruel {
		var figurinha int
		fmt.Scan(&figurinha)
		baruel[i] = figurinha
		baruelFalta[figurinha] = true
		baruelRepetidos[figurinha]++
	}

	printaRepetidos(baruel, baruelRepetidos)
	printaFaltantes(figs, baruelFalta)
}
