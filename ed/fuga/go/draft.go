package main

import "fmt"

func fuga(H, P, F, D int) string {
	for F != H && F != P {
		if D == 1 {
			F = (F + 1) % 16
		}

		if D == -1 {
			F = (F - 1 + 16) % 16
		}
	}

	if F == H {
		return "S"
	}
	return "N"
}

func main() {
	var H, P, F, D int
	fmt.Scan(&H, &P, &F, &D)
	fmt.Println(fuga(H, P, F, D))
}
