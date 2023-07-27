package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	cases := []struct {
		str  string
		want bool
	}{
		{"()(){}[]", true},
		{"()", true},
		{"{{{{{{{[[]]}}}}}}}", true},
		{"{", false},
		{"}", false},
		{"(((([))))]{}", false},
		{"(])",false},
	}

	for _, c := range cases {
		if got := isValid(c.str); got != c.want {
			t.Errorf("isValid('%s') = %t, want %t", c.str, got, c.want)
		}
	}
}
