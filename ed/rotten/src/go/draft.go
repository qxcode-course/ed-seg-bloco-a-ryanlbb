package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	apple := make([][]int, m)

	for i := range m {
		for j := range n {
			fmt.Scan(&apple[i][j])
		}
	}
}
