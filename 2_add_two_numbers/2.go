package main

import "algos/types"

func addTwoNumbers(l1 *types.ListNode, l2 *types.ListNode) *types.ListNode {
	res := &types.ListNode{}
	temp := res

	for l1 != nil || l2 != nil {
		if l1 != nil {
			temp.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp.Val += l2.Val
			l2 = l2.Next
		}
		if temp.Val >= 10 {
			temp.Val -= 10
			temp.Next = &types.ListNode{Val: 1}
		} else if l1 != nil || l2 != nil {
			temp.Next = &types.ListNode{}
		}
		temp = temp.Next
	}

	return res
}
