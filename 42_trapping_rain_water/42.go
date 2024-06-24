package main

func trap(height []int) int {
	m := 0
	for i, v := range height {
		if v > height[m] {
			m = i
		}
	}
	p1 := height[0]
	ans := 0
	for i := 0; i < m; i++ {
		if p1 < height[i] {
			p1 = height[i]
		} else {
			ans += p1 - height[i]
		}
	}

	p2 := height[len(height)-1]
	for i := len(height) - 1; i > m; i-- {
		if p2 < height[i] {
			p2 = height[i]
		} else {
			ans += p2 - height[i]
		}
	}
	return ans
}
