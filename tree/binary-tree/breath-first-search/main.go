package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return [][]int{}
	}

	result = append(result, []int{root.Val})
	if root.Left != nil || root.Right != nil {
		result = append(result, []int{root.Left.Val, root.Right.Val})
		if root.Left != nil {
			result = append(result, levelOrder(root.Left)...)
		}
		if root.Left == nil && root.Right != nil {
			result = append(result, levelOrder(root.Right)...)
		}

	}

	return result
}

func main() {
	input := TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}

	// [[3], [9,2], [3]]
	fmt.Println(levelOrder(&input))
}
