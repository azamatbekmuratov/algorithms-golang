package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := maxDepth(root.Left)
	rightHeight := maxDepth(root.Right)
	max := 1
	if leftHeight > rightHeight {
		max += leftHeight
	} else {
		max += rightHeight
	}
	return max
}

func main() {
	input := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}

	fmt.Println(maxDepth(input))
}
