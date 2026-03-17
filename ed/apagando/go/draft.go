package main
import "fmt"

type Pessoa struct {
    Id  int
    naFila bool
}

func main() {
	var nFila, nDeixouFila int
	fmt.Scan(&nFila)

	fila := make([]Pessoa, nFila)
	
	for i := range fila {
		fmt.Scan(&fila[i].Id)
		fila[i].naFila = true
	}

	fmt.Scan(&nDeixouFila)
	for range nDeixouFila {
		var valor int
		fmt.Scan(&valor)
	}

	for i := 0; i < nFila-1; i++ {
		if fila[i].naFila == true {
			fmt.Print(fila[i].Id, " ")
		}
	}

	fmt.Println()
}
