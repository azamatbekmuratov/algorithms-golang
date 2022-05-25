package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func main() {
	myList := LinkedList{}
	node1 := &Node{data: 33}
	node2 := &Node{data: 12}
	myList.prepend(node1)
	myList.prepend(node2)
	fmt.Println(myList)
}
