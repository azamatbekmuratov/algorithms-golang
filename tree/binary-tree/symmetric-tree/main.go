package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(node1 *TreeNode, node2 *TreeNode) bool {
	// if both trees are empty, then they are mirror image
	if node1 == nil && node2 == nil {
		return true
	}

	// For two trees to be mirror images, the following
	// three conditions must be true
	// 1.) Their root node's key must be same
	// 2.) left subtree of left tree and right subtree
	// of right tree have to be mirror images
	// 3.) right subtree of left tree and left subtree
	// of right tree have to be mirror images
	if node1 != nil && node2 != nil && node1.Val == node2.Val {
		return (isMirror(node1.Left, node2.Right) && isMirror(node1.Right, node2.Left))
	}

	// if none of the above conditions is true then
	// root1 and root2 are not mirror images
	return false
}

func main() {
	input := TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{2, nil, nil}}

	// [[3], [2,2]]
	fmt.Println(isSymmetric(&input))
}
