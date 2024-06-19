package findnumberswithevennumberofdigits

import "testing"

func Test_findNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{12},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{123},
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{12, 345, 2, 6, 7896},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumbers(tt.args.nums); got != tt.want {
				t.Errorf("findNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
