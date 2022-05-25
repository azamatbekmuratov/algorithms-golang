package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack []TreeNode

func (s *Stack) push(input TreeNode) {
	*s = append(*s, input)
}

func (s *Stack) pop() *TreeNode {
	n := len(*s) - 1
	var stack []TreeNode = *s
	poppedElem := stack[n]
	*s = stack[:n]
	return &poppedElem
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack Stack = Stack{}
	var output []int = []int{}
	var curr *TreeNode = root
	/*
	*push root value to the stack
	*check left child
	*do while there is ancesstors of left child
	*if there is child push right then left child to the stack
	*when there is no children
	*pop from the stack left and right child and append to result list
	*
	*if there is no ancesstors check right child
	*
	 */

	for curr != nil {
		fmt.Printf("curr: %v\n", curr)
		stack.push(*curr)
		curr = curr.Left
		if curr.Left != nil {
			stack.push(*curr.Left)
			if curr.Left.Right != nil {
				stack.push(*curr.Left.Right)
			}
			if curr.Left.Left != nil {
				stack.push(*curr.Left.Left)
			}
		}
		curr = curr.Right
		// curr = stack.pop()
		// output = append(output, curr.Val)
		// curr = curr.Right
	}
	return output
}

func main() {
	input := &TreeNode{1, &TreeNode{4, nil, nil}, nil}
	result := postorderTraversal(input)
	fmt.Println(result)
}
