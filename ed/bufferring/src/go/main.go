package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deque struct {
	data     []int
	front    int
	size     int
	capacity int
}

func (dq *Deque) Len() int {
	return dq.size
}

func (dq *Deque) resize(newCap int) {
	newDq := make([]int, newCap)
	for iterations, i := 0, dq.front; iterations < dq.size; iterations, i = iterations+1, (i+1)%dq.capacity {
		newDq[iterations] = dq.data[i]
	}
	dq.data = newDq
	dq.capacity = newCap
	dq.front = 0
}

func findNewCap(cap int) int {
	if cap == 0 {
		return 1
	}
	return cap * 2
}

func (dq *Deque) PushBack(v int) {
	if dq.size == dq.capacity {
		newCap := findNewCap(dq.size)
		dq.resize(newCap)
	}

	dq.data[(dq.front+dq.size)%dq.capacity] = v
	dq.size++
}

func (dq *Deque) PushFront(v int) {
	if dq.size == dq.capacity {
		newCap := findNewCap(dq.size)
		dq.resize(newCap)
	}

	dq.front = (dq.front - 1 + dq.capacity) % dq.capacity
	dq.data[dq.front] = v
	dq.size++
}

func (dq *Deque) PopBack() error {
	if dq.size == 0 {
		return fmt.Errorf("fail: buffer vazio")
	}

	dq.size--
	return nil
}

func (dq *Deque) PopFront() error {
	if dq.size == 0 {
		return fmt.Errorf("fail: buffer vazio")
	}

	dq.front = (dq.front + 1) % dq.capacity
	dq.size--
	return nil
}

func (dq *Deque) Front() (int, error) {
	if dq.size == 0 {
		return 0, fmt.Errorf("fail: buffer vazio")
	}

	return dq.data[dq.front], nil
}

func (dq *Deque) Back() (int, error) {
	if dq.size == 0 {
		return 0, fmt.Errorf("fail: buffer vazio")
	}

	return dq.data[(dq.front+dq.size-1+dq.capacity)%dq.capacity], nil
}

func (dq *Deque) Clear() {
	dq.size = 0
	dq.front = 0
}

func (b *Deque) String() string {
	result := []string{}
	for i := range b.size {
		val := b.data[(b.front+i)%b.capacity]
		result = append(result, fmt.Sprint(val))
	}
	return "[" + strings.Join(result, ", ") + "]"
}

func (b *Deque) Debug() string {
	result := make([]string, b.capacity)
	for i := range result {
		result[i] = " _"
		if i == b.front {
			result[i] = ">_"
		}
	}
	for i := range b.size {
		index := (b.front + i) % b.capacity
		val := b.data[index]
		prefix := " "
		if i == 0 {
			prefix = ">"
		}
		result[index] = fmt.Sprintf("%s%d", prefix, val)
	}
	return strings.Join(result, " |")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := &Deque{data: make([]int, 4), capacity: 4}

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(buf.String())
		case "debug":
			fmt.Println(buf.Debug())
		case "size":
			fmt.Println(buf.Len())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				buf.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				buf.PushFront(num)
			}
		case "pop_back":
			if err := buf.PopBack(); err != nil {
				fmt.Println(err)
			}
		case "pop_front":
			if err := buf.PopFront(); err != nil {
				fmt.Println(err)
			}
		case "front":
			if val, err := buf.Front(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		case "back":
			if val, err := buf.Back(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		case "clear":
			buf.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
