package main

func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}
	dict := make(map[int][2]int, len(trust))

	for _, val := range trust {
		arr := dict[val[1]]
		arr[0] += 1
		dict[val[1]] = arr

		arr2 := dict[val[0]]
		arr2[1] += 1
		dict[val[0]] = arr2
	}

	for k, v := range dict {
		if v[0] == n-1 && v[1] == 0 {
			return k
		}
	}
	return -1
}
