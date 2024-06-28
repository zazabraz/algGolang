package insertion

func sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
			j--
		}

	}
	return arr
}

//0,2,3,1
