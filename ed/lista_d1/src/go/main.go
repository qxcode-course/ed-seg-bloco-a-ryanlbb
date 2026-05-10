package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
}

type LList struct {
	root *Node
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{}

	list.root.next = list.root
	list.root.prev = list.root

	return list
}

func (l *LList) Size() int {
	cont := 0

	// enquanto node nao for root aumenta o contador
	for node := l.root.next; node != l.root; node = node.next {
		cont++
	}

	return cont
}

func (l *LList) Clear() {
	l.root.prev = l.root
	l.root.next = l.root
}

func (l *LList) PushFront(n int) {
	newNode := &Node{Value: n}
	firstAfterRoot := l.root.next

	newNode.prev = l.root
	newNode.next = firstAfterRoot

	l.root.next.prev = newNode
	l.root.next = newNode
}

func (l *LList) PushBack(n int) {
	newNode := &Node{Value: n}
	firstBeforeRoot := l.root.prev

	newNode.prev = firstBeforeRoot
	newNode.next = l.root

	firstBeforeRoot.next = newNode
	l.root.prev = newNode
}

func (l *LList) PopFront() {
	if l.Size() == 0 {
		return
	}

	nodeToPop := l.root.next
	newNext := nodeToPop.next

	l.root.next = newNext
	newNext.prev = l.root

	nodeToPop.next = nil
	nodeToPop.prev = nil
}

func (l *LList) PopBack() {
	if l.Size() == 0 {
		return
	}

	nodeToPop := l.root.prev
	newPrev := nodeToPop.prev

	l.root.prev = newPrev
	newPrev.next = l.root

	nodeToPop.next = nil
	nodeToPop.prev = nil
}

func (l *LList) String() string {
	list := "["
	for node := l.root.next; node != l.root; node = node.next {
		value := fmt.Sprint(node.Value)
		list += value

		if node.next != l.root {
			list += ", "
		}
	}
	list += "]"

	return list
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

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
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
