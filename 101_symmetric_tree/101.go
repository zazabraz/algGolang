package main

import "algos/types"

func isSymmetric(root *types.TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(l, r *types.TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil {
		return false
	}

	if l.Val != r.Val {
		return false
	}

	return isMirror(l.Left, r.Right) && isMirror(l.Right, r.Left)
}
