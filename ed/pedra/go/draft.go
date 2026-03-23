package main

import "fmt"

func diferencaAbs(a, b int) int {
	if a > b {
		return a - b
	}
	return (a - b) * -1
}

func main() {
	var n int
	fmt.Scan(&n)
	diferencas := make([]int, n)

	for i := range n {
		var jgd1, jgd2 int
		fmt.Scan(&jgd1, &jgd2)

		if jgd1 < 10 || jgd2 < 10 {
			diferencas[i] = -1
			continue
		} else {
			diferencas[i] = diferencaAbs(jgd1, jgd2)
		}
	}

	menor := -1
    menorVal := 999999999999999999

	for i := range diferencas {
		if diferencas[i] < menorVal && diferencas[i] != -1 {
			menor = i
		} else {
			continue
		}
	}

	if menor == -1 {
		fmt.Println("sem ganhador")
		return
	}
	fmt.Println(menor)
}
