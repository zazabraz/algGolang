package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, val := range nums {
		fmt.Println(target - val)
		if k, ok := hash[nums[i]]; ok {
			return []int{k,i}
		}
		hash[target-val] = i
	}
	return []int{}
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}
