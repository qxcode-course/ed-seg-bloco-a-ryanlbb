package main

import "fmt"

func ehPrimo(n, div int) bool {
	if n == div {
		return true
	}

	if n <= 1 || n%div == 0 {
		return false
	}

	return ehPrimo(n, div)
}

func main() {
	fmt.Println("Hello, World!")
}
