package main

import "fmt"

func isAnagram(s string, t string) bool {
 
	seen := make(map[rune]int)
	var counter int
	for _, v := range s {
		seen[v]++
		counter++
	}
	for _, v := range t {
		if c, ok := seen[v]; ok && c > 0 {
			seen[v]--
			counter--
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
