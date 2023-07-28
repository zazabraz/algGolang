package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := map[int]struct{}{}
	for _, v := range nums {
		if _, ok := seen[v]; ok {
			return true
		}
		seen[v] = struct{}{}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(containsDuplicate(nums))
}
