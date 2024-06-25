package main

import (
	"algos/types"
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *types.ListNode
		l2 *types.ListNode
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode
	}{
		{
			name: "1",
			args: args{
				l1: &types.ListNode{
					Val: 2,
					Next: &types.ListNode{
						Val: 4,
						Next: &types.ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				l2: &types.ListNode{
					Val: 5,
					Next: &types.ListNode{
						Val: 6,
						Next: &types.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &types.ListNode{
				Val: 2,
				Next: &types.ListNode{
					Val:  4,
					Next: nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
