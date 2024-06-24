package main

import "algos/types"

func deleteMiddle(head *types.ListNode) *types.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var p1, p2 = head, head
	var prev *types.ListNode
	for p2 != nil && p2.Next != nil {
		prev = p1
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	prev.Next = p1.Next

	return head
}
