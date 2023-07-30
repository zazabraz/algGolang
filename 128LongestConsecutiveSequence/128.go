package main

import "fmt"

func longestConsecutive(nums []int) int {
	mp := make(map[int]struct{})
	for _, num := range nums {
		mp[num] = struct{}{}
	}

	ans := 0

	for n := range mp {
		if _, ok := mp[n-1]; !ok {
			temp := 1
			for {
				if _, ok := mp[n+temp]; ok {
					temp++
					continue
				}
				if ans<temp{
					ans=temp
				}
				break
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}

	fmt.Println(longestConsecutive(nums))
}
