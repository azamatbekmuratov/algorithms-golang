package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Levels [][]int

func (l *Levels) append(val, level int) {
	levels := *l
	levels[level] = append(levels[level], val)
	*l = levels
}

func levelOrder(root *TreeNode) [][]int {
	// levels := Levels{}
	levels := [][]int{}
	helper(*root, &levels, 0)
	return levels
}

func helper(node TreeNode, levels *[][]int, level int) {
	//start the current level
	if len(*levels) == level {
		*levels = append(*levels, []int{})
	}

	//fulfil the current level
	levelsLocal := *levels
	levelsLocal[level] = append(levelsLocal[level], node.Val)
	*levels = levelsLocal

	//process child nodes for the next level
	if node.Left != nil {
		helper(*node.Left, levels, level+1)
	}
	if node.Right != nil {
		helper(*node.Right, levels, level+1)
	}

}

func main() {
	input := &TreeNode{4, &TreeNode{5, nil, nil}, &TreeNode{6, nil, nil}}
	result := levelOrder(input)
	fmt.Println(result)
}
