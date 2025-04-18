package nextpermutation

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
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
				nums: []int{4, 2, 0, 2, 3, 2, 0},
			},
			want: []int{4, 2, 0, 3, 0, 2, 2},
		},
		{
			args: args{
				nums: []int{1, 3, 2},
			},
			want: []int{2, 1, 3},
		},
		{
			args: args{
				nums: []int{3, 2, 1},
			},
			want: []int{1, 2, 3},
		},
		{
			args: args{
				nums: []int{1, 2, 1},
			},
			want: []int{2, 1, 1},
		},
		{
			args: args{
				nums: []int{1, 2, 3},
			},
			want: []int{1, 3, 2},
		},
		{
			args: args{
				nums: []int{1, 2},
			},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)

			got := tt.args.nums
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
