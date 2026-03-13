package main
import "fmt"
func main() {
	var fila1, fila2 int
	fmt.Scan(&fila1)
	for _, v := range fila1 {
		var valor int
		fmt.Scan(&valor)
		v = valor
	}
}
