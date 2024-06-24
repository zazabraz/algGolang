package main

func searchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)-1
	for r >= l {
		mid := (r + l) / 2

		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for r >= l {
		mid := (r + l) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}

	return -1
}
