package main

import "fmt"

//nums = [1,2,2,3,3], k = 2

func topKFrequent(nums []int, k int) []int {
	if len(nums) < 1 {
		return nums
	}
	mp := make(map[int]int)
	freq := make([][]int, len(nums)+1)
	//freq:		0 1 2 3 4 5 6
	//nums:		0 0 2 0 0 1 0
	for _, n := range nums {
		mp[n]++
	}
	for k, v := range mp {
		freq[v] = append(freq[v], k)
	}
	ans := []int{}
	for i := len(freq) - 1; i > 0; i-- {
		ans = append(ans, freq[i]...)
		if len(ans) == k {
			return ans
		}
	}
	return ans
}

func main() {
	nums := []int{1, 1, 1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
