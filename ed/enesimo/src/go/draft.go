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

func gerar_enesimo_primo(idx int) int {
    primo := 2
    indice := 1
	
	for indice <= idx {
		if !eh_primo(primo, 2) {
			primo++
			continue
		} else {
			indice++
			continue
		}
	}
	
    return primo
}

func main() {
    var x int
    fmt.Scan(&x)
	fmt.Println(gerar_enesimo_primo(x))
}
