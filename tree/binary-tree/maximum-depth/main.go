package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var leftDepth, rightDepth int = 0, 0
	if root == nil {
		return 0
	}
	_ = leftDepth
	_ = rightDepth
	if root.Left != nil {
		leftDepth = maxDepth(root.Left)
	}
	if root.Right != nil {
		rightDepth = maxDepth(root.Right)
	}

	return findMax(leftDepth, rightDepth) + 1
}

func findMax(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func main() {
	input := TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}

	// [[3], [9,2], [3]]
	fmt.Println(maxDepth(&input))
}
