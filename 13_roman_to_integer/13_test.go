package main

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, c := range cases {
		if got := romanToInt(c.in); got != c.want {
			t.Errorf("romanToInt(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}
