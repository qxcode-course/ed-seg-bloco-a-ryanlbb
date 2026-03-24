package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	calmMen := make([]int, 0)
	for _, v := range vet {
		if v > 0 {
			calmMen = append(calmMen, v)
		}
	}
	return calmMen
}

func getCalmWomen(vet []int) []int {
	calmWomen := make([]int, 0)
	for _, v := range vet {
		if v < 0  && v > -10 {
			calmWomen = append(calmWomen, v)
		}
	}
	return calmWomen
}

func sortVet(vet []int) []int {
	_ = vet
	return nil
}

func sortStress(vet []int) []int {
	_ = vet
	return nil
}

func reverse(vet []int) []int {
	reverse := make([]int, 0)
	for i := len(vet) - 1; i >= 0; i-- {
		reverse = append(reverse, vet[i])
	}
	return reverse
}

func unique(vet []int) []int {
	mapUnicos := make(map[int]int)
	semrep := make([]int, 0)
	for i := range vet {
		if mapUnicos[vet[i]] < 1 {
			mapUnicos[vet[i]]++
			semrep = append(semrep, vet[i])
		}
	}

	for i := range semrep {
		if i < len(semrep)-1 {
			return semrep[i]
		}
	}
	return nil
}

func repeated(vet []int) []int {
	_ = vet
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

