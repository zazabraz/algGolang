package main

import "fmt"

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	var mid int
	for low <= high {
		mid = (low + high) / 2
		switch {
		case nums[mid] < target:
			low = mid + 1
		case nums[mid] > target:
			high = mid - 1
		default:
			return mid
		}
	}
	return low
}

func main() {
	nums := []int{1, 3, 5, 6, 8, 9, 11}
	fmt.Println(searchInsert(nums, 2))
}
