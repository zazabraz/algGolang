package main

import "strings"

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for _, str := range strs {
		pre := make([]string, 0)
		for i, _ := range str {
			if i > len(prefix)-1 {
				break
			} else {
				if str[i] == prefix[i] {
					pre = append(pre, string(str[i]))
				} else {
					break
				}
			}
		}
		prefix = strings.Join(pre, "")
	}
	return prefix
}
