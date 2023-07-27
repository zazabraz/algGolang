package main

import "fmt"

func romanToInt(s string) int {
	romans := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	ans := romans[s[len(s)-1]]
	for i := len(s) - 1; i > 0; i-- {
		if romans[s[i]] > romans[s[i-1]] {
			ans -= romans[s[i-1]]
		} else {
			ans += romans[s[i-1]]
		}
	}
	return ans
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
