package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head *ListNode = &ListNode{}
	if list1.Val < list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}

	head.Next = mergeTwoLists(list1, list2)
	return head
}

func main() {

	node3List1 := &ListNode{Val: 4, Next: nil}
	node2List1 := &ListNode{Val: 2, Next: node3List1}
	node1List1 := &ListNode{Val: 1, Next: node2List1}

	node3List2 := &ListNode{Val: 6, Next: nil}
	node2List2 := &ListNode{Val: 3, Next: node3List2}
	node1List2 := &ListNode{Val: 1, Next: node2List2}

	fmt.Printf("Result: %#v", mergeTwoLists(node1List1, node1List2))
}
