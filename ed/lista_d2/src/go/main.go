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

	root *Node
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}

func (ll *LList) Front() *Node {
	if ll.root.Next() == nil {
		return nil
	}
	return ll.root.next
}

func (ll *LList) Back() *Node {
	if ll.root.Prev() == nil {
		return nil
	}
	return ll.root.prev
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	ll := &LList{
		root: &Node{Value: 0},
		size: 0,
	}

	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.root.root = ll.root

	return ll
}

func (ll *LList) Size() int {
	return ll.size
}

func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}

func (ll *LList) PushFront(value int) {
	newNode := &Node{
		Value: value,
		next:  ll.root.next,
		prev:  ll.root,
		root:  ll.root,
	}

	ll.root.next.prev = newNode
	ll.root.next = newNode

	ll.size++
}

func (ll *LList) PushBack(value int) {
	newNode := &Node{
		Value: value,
		next:  ll.root,
		prev:  ll.root.prev,
		root:  ll.root,
	}

	ll.root.prev.next = newNode
	ll.root.prev = newNode

	ll.size++
}

func (ll *LList) PopFront() {
	newRootNext := ll.root.next.next
	PoppedNode := ll.root.next

	PoppedNode.prev = nil
	ll.root.next = newRootNext
	PoppedNode.next = nil
}

func (ll *LList) PopBack() {
	newRootPrev := ll.root.prev.prev
	PoppedNode := ll.root.prev

	PoppedNode.prev = nil
	ll.root.prev = newRootPrev
	PoppedNode.prev = nil
}

func (ll *LList) Search(value int) *Node {
	for node := ll.root.Next(); node != nil; node = node.Next() {
		if value == node.Value {
			return node
		}
	}
	return nil
}

func (ll *LList) Insert(node *Node, value int) {
	if node == nil {
		return
	}

	newNode := &Node{
		Value: value,
		next:  node,
		prev:  node.prev,
		root:  node.root,
	}

	oldPrevious := node.prev
	node.prev = newNode
	oldPrevious.next = newNode
	ll.size++
}

func (ll *LList) Remove(node *Node) *Node {
	newRefNode := node.next

	node.prev.next = newRefNode
	node.prev = nil
	node.next = nil

	ll.size--

	return newRefNode
}

func (ll *LList) String() {
	fmt.Print("[")
	for node := ll.root.Next(); node != nil; node = node.Next() {
		fmt.Print(node.Value)

		if node != ll.Back() {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
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
			ll.String()
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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
