package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return []int{}
	}
	if root.Left != nil {
		result = append(result, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		result = append(result, postorderTraversal(root.Right)...)
	}
	result = append(result, root.Val)

	return result
}

func main() {
	input := TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}

	fmt.Println(postorderTraversal(&input))
}
