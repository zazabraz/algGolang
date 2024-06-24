package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	mp := make(map[[26]int][]string)
	
	for _, s := range strs {
		k := [26]int{}
		for _, ch := range s {
			k[int(ch)-'a'] += 1
		}
		mp[k] = append(mp[k], s)
	}
	res := [][]string{}
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ans := groupAnagrams(strs)
	for _, v := range ans {
		fmt.Println(v)
	}
}
