package main

func maxArea(height []int) int {
	res := 0
	l, r := 0, len(height)-1
	for l != r {
		min := 0
		if height[l] < height[r] {
			min = height[l]
		} else {
			min = height[r]
		}

		lenn := r - l
		if res < lenn*min {
			res = lenn * min
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}
