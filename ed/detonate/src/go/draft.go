package main

import "fmt"

fun detonate(grid [][]int) {

}

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	bombs := make([][]int, m)
	for i := range m {
		bombs[i] = make([]int, n)
	}

	for l := range bombs {
		for c := range bombs[l] {
			fmt.Scan(&bombs[l][c])
		}
	}

}

