package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev ListNode
	var curr = head
	for curr != nil {
		var nextTemp ListNode = *curr.Next
		curr.Next = &prev
		prev = *curr
		curr = &nextTemp
	}
	return &prev

}

func main() {

	var node4 ListNode = ListNode{Val: 5, Next: nil}
	var node3 ListNode = ListNode{Val: 4, Next: &node4}
	var node2 ListNode = ListNode{Val: 3, Next: &node3}
	var node1 ListNode = ListNode{Val: 2, Next: &node2}

	reverseList(&node1)
	// fmt.Println(node1)
}
