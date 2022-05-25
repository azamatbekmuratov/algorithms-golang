package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	firstNode := head
	secondNode := head.Next

	//Swapping
	firstNode.Next = swapPairs(secondNode.Next)
	secondNode.Next = firstNode
	return secondNode
}

func main() {
	var node4 ListNode = ListNode{Val: 4, Next: nil}
	var node3 ListNode = ListNode{Val: 3, Next: &node4}
	var node2 ListNode = ListNode{Val: 2, Next: &node3}
	var node1 ListNode = ListNode{Val: 1, Next: &node2}
	fmt.Println(swapPairs(&node1))

	fmt.Println(node1)
	fmt.Println(node2)
	fmt.Println(node3)
	fmt.Println(node4)
}
