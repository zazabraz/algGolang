package main

//TODO
func removeElement(nums []int, val int) int {
	res := make([]int, 0)
	//delete pointer
	p1 := 0
	//findPointer
	p2 := 0

	for p1 != len(nums) && p2 != len(nums) {
		if nums[p1] == val {
			for i := p2; ; i++ {
				if nums[i] != val {
					res = append(res, nums[i])
					p2 = i + 1
					break
				}
			}
		}
		p1++
	}

	return len(res)
}
