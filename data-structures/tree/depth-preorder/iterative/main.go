package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack []TreeNode

func (s *Stack) push(input TreeNode) {
	*s = append(*s, input)
}

func (s *Stack) pop() TreeNode {
	var n int = len(*s) - 1 //Top element
	var stack []TreeNode = *s
	*s = stack[:n] // Pop
	fmt.Printf("pop value: %v\n", stack[n])
	return stack[n]
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack = Stack{}
	var output = []int{}
	stack.push(*root)
	for !stack.isEmpty() {
		var node TreeNode = stack.pop()
		output = append(output, node.Val)
		if node.Right != nil {
			stack.push(*node.Right)
		}
		if node.Left != nil {
			stack.push(*node.Left)
		}
	}

	return output
}

func main() {
	input := &TreeNode{4, &TreeNode{5, nil, nil}, nil}
	result := preorderTraversal(input)
	fmt.Println(result)
}
