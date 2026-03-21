package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	var p float64 = (a + b + c) / 2
	var area float64 = math.Sqrt(p * (p - a) * (p - b) * (p - c))

	fmt.Printf("%.2f\n", area)
}
