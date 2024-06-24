package main

import "algos/types"

func middleNode(head *types.ListNode) *types.ListNode {
	var p1, p2 = head, head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return p1
}
