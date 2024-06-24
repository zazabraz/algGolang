package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	p1, p2 := 0, len(s)-1
	for p1 < p2 {
		l := rune(s[p1])
		r := rune(s[p2])
		if !unicode.IsLetter(l) && !unicode.IsNumber(l) {
			p1++
		} else if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			p2--
		} else if strings.EqualFold(string(s[p1]), string(s[p2])) {
			p1++
			p2--
		} else {
			return false
		}
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	ans := isPalindrome(s)
	fmt.Println(ans)
}
