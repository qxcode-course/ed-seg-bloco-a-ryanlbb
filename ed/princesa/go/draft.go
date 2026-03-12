package main

import (
	"fmt"
)

func main() {
	var N, espada int
	fmt.Scan(&N, &espada)

	vitima := espada
	fila := make([]int, N)

	for i := range fila {
		fila[i] = i + 1
	}

	mortos := 0

    //mostra todos
    fmt.Print("[ ")
	for _, v := range fila {
	    if espada == v{
			fmt.Printf("%d> ", v)
		} else {
            fmt.Printf("%d ", v)
        }
	}
    fmt.Println("]")
    fila[vitima] = 0
    mortos++
		
    
    //matança
	for mortos < N{ 
		for _, v := range fila {
            fmt.Print("[ ")

			if espada == v {
				fmt.Printf("%d> ", v)
			} else {
                fmt.Printf("%d ", v)
            }

            mortos++
		    fmt.Println("]")
        }
        
	}
}
