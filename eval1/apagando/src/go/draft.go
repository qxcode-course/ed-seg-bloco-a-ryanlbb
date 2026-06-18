package main

import "fmt"

func main() {
	var nVet, nApagando int
	fmt.Scan(&nVet)
	vet := make([]int, nVet)
	mapApagando := make(map[int]bool)
	for i := range vet {
		fmt.Scan(&vet[i])
		mapApagando[vet[i]] = true
	}

	fmt.Scan(&nApagando)
	vetApagar := make([]int, nApagando)
	for i := range nApagando {
		fmt.Scan(&vetApagar[i])
		mapApagando[vetApagar[i]] = false
	}

	vetFinal := make([]int, 0)
	for _, v := range vet {
		if mapApagando[v] == true {
			vetFinal = append(vetFinal, v)
		}
	}

	for i, v := range vetFinal {
		fmt.Print(v)

		if i <= len(vetFinal)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
