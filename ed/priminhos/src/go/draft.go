package main

import "fmt"

func isPrime(n, div int) bool {
	if n == div {
		return true
	}

	if n < 2 || n%div == 0 {
		return false
	}

	return isPrime(n, div+1)
}

func listPrime(vet []int, n, cont, currentNum int) {
	if n == cont {
		toStr(vet)
		return
	}

	if isPrime(currentNum, 2) {
		cont++
		vet = append(vet, currentNum)
	}

	listPrime(vet, n, cont, currentNum+1)
}

func toStr(vet []int) {
	fmt.Print("[")
	for i, v := range vet {
		fmt.Print(v)
		if i != len(vet)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func main() {
	var n int
	fmt.Scan(&n)
	vet := make([]int, 0, n)
	listPrime(vet, n, 0, 0)
}
