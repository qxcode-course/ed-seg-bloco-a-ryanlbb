package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Multiset struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *Multiset {
	newMS := make([]int, capacity)
	return &Multiset{
		data:     newMS,
		size:     0,
		capacity: capacity,
	}
}

func (ms *Multiset) expand() {
	newCap := ms.size * 2

	if newCap == 0 {
		newMS := make([]int, 1)
		ms.data = newMS
		ms.capacity = 1
	} else {
		newMS := make([]int, newCap)
		copy(newMS, ms.data)
		ms.data = newMS
		ms.capacity = newCap
	}
}

func (ms *Multiset) findInsertionIdx(value int) int {
	l := 0
	r := ms.size - 1

	for l <= r {
		mid := (l + r) / 2

		if ms.data[mid] < value {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}

func (ms *Multiset) findValueIdx(value int) int {
	l := 0
	r := ms.size - 1

	for l <= r {
		mid := (l + r) / 2

		if ms.data[mid] == value {
			return mid
		}

		if ms.data[mid] < value {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

func (ms *Multiset) Erase(value int) error {
	if ms.findValueIdx(value) == -1 {
		return fmt.Errorf("value not found")
	}

	ms.erase(ms.findValueIdx(value))
	return nil
}

func (ms *Multiset) erase(index int) {
	for i := index; i < ms.size; i++ {
		ms.data[i] = ms.data[i+1]
	}
	ms.size--
}

func (ms *Multiset) Insert(value int) {
	if ms.size == ms.capacity {
		ms.expand()
	}

	ms.insert(value, ms.findInsertionIdx(value))
}

func (ms *Multiset) insert(value, index int) {
	for i := ms.size; i > index; i-- {
		ms.data[i] = ms.data[i-1]
	}
	ms.data[index] = value
	ms.size++
}

func (ms *Multiset) Contains(value int) bool {
	return ms.findValueIdx(value) != -1
}

func (ms *Multiset) Count(value int) int {
	firstOcurrence := ms.findInsertionIdx(value)
	count := 0

	for i := firstOcurrence; i < ms.size && ms.data[i] == value; i++ {
		count++
	}

	return count
}

func (ms *Multiset) Unique() int {
	if ms.size == 0 {
		return 0
	}
	count := 1

	for i := range ms.size - 1 {
		if ms.data[i] != ms.data[i+1] {
			count++
		}
	}

	return count
}

func (ms *Multiset) Clear() {
	ms.size = 0
}

func (ms *Multiset) String() string {
	return "[" + Join(ms.data[:ms.size], ", ") + "]"
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
