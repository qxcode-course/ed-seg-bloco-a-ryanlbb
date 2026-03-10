package main

import (
	"fmt"
)
func main() {
    var N, E int
    fmt.Scan(&N, &E)

    fila := make([]int, N)

    for i := range fila {
        fila[i] = i + 1
    }

    
    fmt.Println(fila)
}
