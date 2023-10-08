package array

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{0, 1},
			},
			want: []int{1, 0},
		},
		{
			args: args{
				nums: []int{0, 0, 1},
			},
			want: []int{1, 0, 0},
		},
		{
			args: args{
				nums: []int{0, 0, 1, 2},
			},
			want: []int{1, 2, 0, 0},
		},
		{
			args: args{
				nums: []int{0, 1, 0, 2},
			},
			want: []int{1, 2, 0, 0},
		},
		{
			args: args{
				nums: []int{ 1 },
			},
			want: []int{ 1 },
		},
		{
			args: args{
				nums: []int{ 0, 1, 0, 2, 0 },
			},
			want: []int{ 1, 2, 0, 0, 0 },
		},
		{
			args: args{
				nums: []int{1, 0, 0, 2},
			},
			want: []int{1, 2, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
