package main

import "fmt"

func eh_primo(x int, div int) bool {
	if x == div {
		return true
	}

	if x <= 1 || x % div == 0 {
		return false
	} 

	return eh_primo(x, div + 1)
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(eh_primo(x, 2))
}
