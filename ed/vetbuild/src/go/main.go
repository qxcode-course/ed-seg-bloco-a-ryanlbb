package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

func (vet *Vector) Status() string {
	status := fmt.Sprintf("size:%d capacity:%d", vet.size, vet.capacity)
	return status
}

func (vet *Vector) Show() string {
	show := fmt.Sprintf("[%s]", Join(vet.data[:vet.size], ", "))
	return show
}

func (vet *Vector) Reserve(newCapacity int) {
	vetDataNew := make([]int, newCapacity)

	for i := range vet.size {
		vetDataNew[i] = vet.data[i]
	}

	vet.data = vetDataNew
	vet.capacity = newCapacity
}

func (vet *Vector) PushBack(x int) {
	if vet.size+1 > vet.capacity {
		if vet.capacity == 0 {
			vet.Reserve(1)
		} else {
			vet.Reserve(vet.capacity * 2)
		}
	}

	vet.data[vet.size] = x
	vet.size++
}

func (vet *Vector) AtRange(idx int) bool {
	return idx >= 0 && idx < vet.size
}

func (vet *Vector) IsEmpty() bool {
	return vet.size <= 0
}

func (vet *Vector) At(idx int) (int, error) {
	if !vet.AtRange(idx) {
		return 0, fmt.Errorf("index out of range")
	}

	return vet.data[idx], nil
}

func (vet *Vector) Set(idx, value int) error {
	if !vet.AtRange(idx) {
		return fmt.Errorf("index out of range")
	}

	vet.data[idx] = value
	return nil
}

func (vet *Vector) Clear() {
	for i := range vet.data {
		vet.data[i] = 0
	}

	vet.size = 0
}

func (vet *Vector) PopBack() error {
	if vet.IsEmpty() {
		return fmt.Errorf("vector is empty")
	}

	vet.data[vet.size-1] = 0
	vet.size--
	return nil
}

func (vet *Vector) Insert(idx, value int) error {
	if vet.size+1 > vet.capacity {
		vet.Reserve(vet.size * 2)
	}
	if !vet.AtRange(idx) {
		return fmt.Errorf("index out of range")
	}

	// determina ate qual indice ele vai andar do fim do vet pro comeco
	for i := vet.size - 1; i >= idx; i-- {
		vet.data[i+1] = vet.data[i]
	}

	vet.data[idx] = value
	vet.size++
	return nil
}

func (vet *Vector) Erase(idx int) error {
	if !vet.AtRange(idx) {
		return fmt.Errorf("index out of range")
	}

	for i := idx; i < vet.size; i++ {
		// se for o ultimo item, nao acessa tamanho idx invalido
		if i == vet.size-1 {
			vet.data[i] = 0
			continue
		}

		vet.data[i] = vet.data[i+1]
	}

	vet.size--
	return nil
}

func (vet *Vector) IndexOf(integer int) int {
	exists := false
	idx := 0
	for i, v := range vet.data {
		if v == integer {
			exists = true
			idx = i
			break
		}
	}

	if !exists {
		return -1
	}
	return idx
}

func (vet *Vector) Contains(integer int) bool {
	return vet.IndexOf(integer) != -1
}

func (vet *Vector) Slice(start, end int) *Vector {
	start = ((start % vet.size) + vet.size) % vet.size
	end = ((end % vet.size) + vet.size) % vet.size

	vetDataNew := Vector{
		data:     vet.data[start:end],
		size:     end - start,
		capacity: vet.capacity,
	}

	return &vetDataNew
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v.Show())
		case "status":
			fmt.Println(v.Status())
		case "pop":
			err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			// fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(slice.Show())
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
