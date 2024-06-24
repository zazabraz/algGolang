package main

import "fmt"

func isAnagram(s string, t string) bool {

	seen := make(map[rune]int)
	for _, v := range s {
		seen[v]++
	}
	for _, v := range t {
		if c, ok := seen[v]; ok && c > 0 {
			seen[v]--
		} else {
			return false
		}
	}
	return true
}

func main() {
	s := "a"
	t := "ab"
	fmt.Println(isAnagram(s, t))
}
