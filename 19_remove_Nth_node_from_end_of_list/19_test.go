package main

import (
	"algos/types"
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *types.ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode
	}{
		{
			name: "1",
			args: args{
				head: &types.ListNode{
					Val: 1,
					Next: &types.ListNode{
						Val:  2,
						Next: nil},
				},
				n: 1,
			},
			want: &types.ListNode{
				Val:  1,
				Next: nil,
			},
		},
		{
			name: "2",
			args: args{
				head: &types.ListNode{
					Val: 1,
					Next: &types.ListNode{
						Val:  2,
						Next: nil},
				},
				n: 2,
			},
			want: &types.ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
