package main

import "fmt"

func main() {
	var totalFig, nBaruelFig int
	fmt.Scan(&totalFig)
	figs := make([]int, totalFig)
	figsBaruel := make(map[int]int)
	repetidas := make([]int, 0)

	// todas as fig
	for i := range totalFig {
		figs[i] = i + 1
	}

	fmt.Scan(&nBaruelFig)

	// fig que o caba tem
	for range nBaruelFig {
		var a int
		fmt.Scan(&a)
		figsBaruel[a]++

		if figsBaruel[a] > 1 {
			repetidas = append(repetidas, a)
		}
	}

	faltas := 0
	for i, v := range figs {
		if figsBaruel[v] > 0 {
			figs[i] = 0
		}
	}

	if len(repetidas) == 0 {
		fmt.Print("N")
	} else {
		for _, v := range repetidas {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()

	if faltas > 0 {
		for _, v := range figs {
			if v != 0 {
				fmt.Print(v, " ")
			}
		}
	} else {
		fmt.Print("N")
	}
	fmt.Println()
	fmt.Println(faltas)
}

