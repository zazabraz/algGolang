package main

func productExceptSelf(nums []int) []int {
	res := make([]int, 0, len(nums))

	counter := 1
	for i := 0; i < len(nums); i++ {
		if i-1 >= 0 {
			counter = counter * nums[i-1]
		}
		res = append(res, counter)
	}

	counter = 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+1 < len(nums) {
			counter = counter * nums[i+1]
		}
		res[i] = res[i] * counter
	}

	return res
}
