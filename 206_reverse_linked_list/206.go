package main

import "algos/types"

func reverseList(head *types.ListNode) *types.ListNode {
	var prev, curr *types.ListNode = nil, head
	for curr != nil {
		next := curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}
	return prev
}

//1 2 3
// 1 iteration
// prev = 1
// curr = 2
//
