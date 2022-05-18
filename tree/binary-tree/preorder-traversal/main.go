package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return []int{}
	}

	result = append(result, root.Val)
	if root.Left != nil {
		result = append(result, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		result = append(result, preorderTraversal(root.Right)...)
	}

	return result
}

func main() {
	input := TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}

	fmt.Println(preorderTraversal(&input))
}
