package main

import (
	"algos/types"
	"fmt"
)

func maxDepth(root *types.TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Left == nil && root.Right == nil {
		return 1
	} else {
		res++
	}
	resL := 0
	if root.Left != nil {
		resL += maxDepth(root.Left)
	}
	resR := 0
	if root.Right != nil {
		resR += maxDepth(root.Right)
	}
	if resR >= resL {
		res += resR
	} else {
		res += resL
	}
	return res
}

func main() {
	tr := &types.TreeNode{
		Val:   3,
		Left:  &types.TreeNode{Val: 9},
		Right: &types.TreeNode{Val: 20, Left: &types.TreeNode{Val: 15}, Right: &types.TreeNode{Val: 7}},
	}
	fmt.Println(maxDepth(tr))
}
