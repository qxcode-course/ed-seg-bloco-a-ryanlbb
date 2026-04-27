package main

import "fmt"

func reproduzir(meses, pares int) int {
	if meses <= 2 {
		return 1
	}

	return reproduzir(meses-1, pares) + pares*reproduzir(meses-2, pares)
}

func main() {
	var meses, pares int
	fmt.Scan(&meses, &pares)
	fmt.Println(reproduzir(meses, pares))
}
