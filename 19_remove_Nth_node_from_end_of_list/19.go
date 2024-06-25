package main

import "algos/types"

func removeNthFromEnd(head *types.ListNode, n int) *types.ListNode {
	if head.Next == nil {
		return nil
	}
	var prev, delPointer *types.ListNode = nil, nil
	curr := head

	for i := 1; curr != nil; i++ {
		if delPointer != nil {
			delPointer = delPointer.Next
		}
		if delPointer == head.Next {
			prev = head
		} else if prev != nil {
			prev = prev.Next
		}
		if i == n {
			delPointer = head
		}
		curr = curr.Next
	}
	if prev != nil {
		prev.Next = delPointer.Next
	} else {
		head = head.Next
	}
	return head
}
