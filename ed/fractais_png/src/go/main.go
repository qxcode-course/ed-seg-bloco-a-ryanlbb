package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func drawTree(pen *Pen, comprimento float64) {
	if comprimento < 2 {
		return
	}
}

func main() {
	pen := NewPen(1000, 1000)
	pen.SetRGB(70, 70, 10)
	pen.SetPosition(500, 1000)
	pen.SetHeading(90)
	pen.Walk(150)

	drawTree(pen, 200)
	pen.SetRGB(255, 0, 0)

	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
