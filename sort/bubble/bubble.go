package bubble

func sort(arr []int) []int {
	for i := range arr {
		for j := 0; j < len(arr)-i; j++ {
			if len(arr) == j+1 {
				continue
			}
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
