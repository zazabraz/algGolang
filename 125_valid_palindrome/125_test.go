package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "1",
			want: true,
			arg:  "A man, a plan, a canal: Panama",
		},
		{
			name: "2",
			want: false,
			arg:  "race a car",
		},
		{
			name: "3",
			want: true,
			arg:  " ",
		},
		{
			name: "4",
			want: true,
			arg:  "aa",
		},
		{
			name: "5",
			want: false,
			arg:  "0P",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.arg); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
