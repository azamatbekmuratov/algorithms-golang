package main

import "fmt"

type TreeNode struct {
	Val   int
	left  *TreeNode
	right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == val {
		return root
	} else {
		if val < root.Val {
			root = searchBST(root.left, val)
		} else {
			root = searchBST(root.right, val)
		}
	}

	return root
}

func main() {
	var treenode = TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{7, nil, nil}}
	fmt.Println(searchBST(&treenode, 2))
}
