package bubble

import (
	"reflect"
	"testing"
)

func Test_sort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{arr: []int{1, 5, 7, 3, 0}},
			want: []int{0, 1, 3, 5, 7},
		},
		{
			name: "2",
			args: args{arr: []int{0, 1}},
			want: []int{0, 1},
		},
		{
			name: "3",
			args: args{arr: []int{1, 0}},
			want: []int{0, 1},
		},
		{
			name: "4",
			args: args{arr: []int{1}},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
