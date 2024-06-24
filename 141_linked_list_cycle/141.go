package main

import "algos/types"

func hasCycle(head *types.ListNode) bool {
	mp := make(map[*types.ListNode]struct{})
	for head != nil {
		if _, ok := mp[head]; ok {
			return true
		}
		mp[head] = struct{}{}
	}
	return false
}
