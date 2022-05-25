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

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack Stack = Stack{}
	var output []int = []int{}
	var curr *TreeNode = root
	for curr != nil || !stack.isEmpty() {
		for curr != nil {
			stack.push(*curr)
			curr = curr.Left
		}
		curr = stack.pop()
		output = append(output, curr.Val)
		curr = curr.Right
	}
	return output
}

func main() {
	input := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	result := inorderTraversal(input)
	fmt.Println(result)
}
