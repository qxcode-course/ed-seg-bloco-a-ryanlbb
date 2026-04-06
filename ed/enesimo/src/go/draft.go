package main

import "fmt"

func eh_primo(x int, div int) bool {
	if x == div {
		return true
	}

	if x <= 1 || x%div == 0 {
		return false
	}

	return eh_primo(x, div+1)
}

func gerar_enesimo_primo(idx, idxAtual, primo int) int {
	if primo > 2 && eh_primo(primo, 2) {
		idxAtual++
	}

	if idx == idxAtual {
		return primo
	}

	return gerar_enesimo_primo(idx, idxAtual, primo + 1)
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(gerar_enesimo_primo(x, 1, 2))
}
