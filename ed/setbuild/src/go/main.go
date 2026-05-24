package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(cap int) *Set {
	return &Set{
		data:     make([]int, cap),
		size:     0,
		capacity: cap,
	}
}

func (s *Set) reserve(newCapacity int) {
	newSet := make([]int, newCapacity)

	copy(newSet, s.data)

	s.capacity = newCapacity
	s.data = newSet
}

func (s *Set) binarySearch(value int) int {
	l := 0
	r := s.size - 1

	for l <= r {
		mid := (l + r) / 2

		if s.data[mid] == value {
			return mid
		}

		if s.data[mid] > value { // se meio maior, direita antes do meio
			r = mid - 1
		} else { // se meio menor, esquerda depois do meio
			l = mid + 1
		}
	}
	return -1
}

func (s *Set) findInsertionIndex(value int) int {
	l := 0
	r := s.size - 1
	for l <= r {
		mid := (l + r) / 2
		if s.data[mid] > value {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func (s *Set) insert(value, index int) {
	if s.size == s.capacity {
		newCap := s.capacity * 2
		if newCap == 0 {
			newCap = 1
		}
		s.reserve(newCap)
	}

	for i := s.size; i > index; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[index] = value
	s.size++
}

func (s *Set) erase(index int) {
	for i := index; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}
	s.size--
}

func (s *Set) Insert(value int) {
	if s.binarySearch(value) != -1 {
		return
	}

	s.insert(value, s.findInsertionIndex(value))
}

func (s *Set) Contains(value int) bool {
	return s.binarySearch(value) != -1
}

func (s *Set) Erase(value int) error {
	if s.binarySearch(value) == -1 {
		return fmt.Errorf("value not found")
	}

	s.erase(s.binarySearch(value))
	return nil
}

func (s *Set) String() string {
	set := "["
	for i := range s.size {
		if i == s.size-1 {
			set += fmt.Sprintf("%d", s.data[i])
			break
		}

		set += fmt.Sprintf("%d, ", s.data[i])
	}
	set += "]"

	return set
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
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
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v.String())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			err := v.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(v.Contains(value))
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
