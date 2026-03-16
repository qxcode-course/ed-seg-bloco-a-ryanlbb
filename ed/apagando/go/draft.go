package main
import "fmt"
func main() {
	var nFila, nDeixouFila int
	fmt.Scan(&nFila)

	fila := make([][2]int, nFila)
	
	for i := range nFila {
		var valor int
		fmt.Scan(&valor)
		fila[i][0] = valor
		fila[i][1] = 1 
	}

	fmt.Scan(&nDeixouFila)
	for range nDeixouFila {
		var valor int
		fmt.Scan(&valor)
		fila[valor][1] = 0
	}

	for i := range nFila {
		if fila[i][1] == 1 {
			fmt.Print(i, " ")
		}
	}

	fmt.Println()
}
