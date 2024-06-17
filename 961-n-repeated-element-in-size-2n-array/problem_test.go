package nrepeatedelementinsize2narray

import "testing"

func Test_repeatedNTimes(t *testing.T) {
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
				nums: []int{1, 2, 3, 3},
			},
			want: 3,
		},
		{
			args: args{
				nums: []int{1, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedNTimes(tt.args.nums); got != tt.want {
				t.Errorf("repeatedNTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
