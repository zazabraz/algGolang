package main

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{
			name: "1",
			strs: []string{"cir", "car"},
			want: "c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
