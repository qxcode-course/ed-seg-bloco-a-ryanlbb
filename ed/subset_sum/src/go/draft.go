package main

import "fmt"

func sumSubset(vet []int, sum int) bool {
	return checkSubset(vet, 0, sum, 0, 0)
}

func checkSubset(vet []int, currentSum, sum, idx, cont int) bool {
	if currentSum == sum {
		return true
	}
	if currentSum > sum || cont == len(vet) {
		return false
	}

	nextIdx := (idx + 1) % len(vet)
	return checkSubset(vet, currentSum+vet[idx], sum, nextIdx, cont+1) ||
		checkSubset(vet, currentSum, sum, nextIdx, cont+1)
}

func main() {
	var n, sum int
	fmt.Scan(&n, &sum)
	vet := make([]int, n)
	for i := range n {
		fmt.Scan(&vet[i])
	}
	fmt.Println(sumSubset(vet, sum))
}

