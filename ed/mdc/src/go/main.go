package main

import (
	"fmt"
)

func encontrarMdc(a, b int) int{
	if a == 0 && b != 0 {
		return b
	}
	if b == 0 && a != 0 {
		return a
	}

	return encontrarMdc(b, a % b)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(encontrarMdc(a, b))
}
