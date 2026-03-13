package main
import "fmt"
func main() {
	var nF1, nF2 int
	fmt.Scan(&nF1)

	fila1 := make([]int, nF1)
	
	for i := range fila1 {
		var valor int
		fmt.Scan(&valor)
		fila1[i] = valor 
	}
}
