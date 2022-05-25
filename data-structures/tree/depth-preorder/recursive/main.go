package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result = []int{}
	result = append(result, root.Val)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}

func main() {
	input := &TreeNode{4, &TreeNode{5, nil, nil}, nil}
	result := preorderTraversal(input)
	fmt.Println(result)
}
