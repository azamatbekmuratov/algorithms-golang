package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue []TreeNode

func (q *Queue) push(node TreeNode) {
	*q = append(*q, node)
}

func (q *Queue) pop() *TreeNode {
	var queue []TreeNode = *q
	*q = queue[1:]
	return &queue[0]
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queue Queue = Queue{}
	var output [][]int = [][]int{}
	var curr TreeNode = *root
	queue.push(curr)
	for !queue.isEmpty() {
		var node TreeNode = *queue.pop()
		fmt.Println(node)
		levelOutput := []int{node.Val}
		output = append(output, levelOutput)
		var children []int = []int{}
		if node.Left != nil {
			queue.push(*node.Left)
		}
		if node.Right != nil {
			queue.push(*node.Right)
		}

	}

	return output
}

func main() {
	input := &TreeNode{4, &TreeNode{5, nil, nil}, &TreeNode{6, nil, nil}}
	result := levelOrder(input)
	fmt.Println(result)
}
