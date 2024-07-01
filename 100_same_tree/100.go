package main

import "algos/types"

func isSameTree(p *types.TreeNode, q *types.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	l := isSameTree(p.Left, q.Left)
	r := isSameTree(p.Right, q.Right)
	return l && r
}
