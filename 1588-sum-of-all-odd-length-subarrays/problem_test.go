package sumofalloddlengthsubarrays

import "testing"

func Test_sumOddLengthSubarrays(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				arr: []int{1, 1},
			},
			want: 2,
		},
		{
			args: args{
				arr: []int{1, 1, 1},
			},
			want: 6,
		},
		{
			args: args{
				arr: []int{1, 4, 2, 5, 3},
			},
			want: 58,
		},
		{
			args: args{
				arr: []int{10, 11, 12},
			},
			want: 66,
		},
		{
			args: args{
				arr: []int{1, 1, 1, 1, 1, 1},
			},
			want: 28,
		},
		{
			args: args{
				arr: []int{6, 9, 14, 5, 3, 8, 7, 12, 13, 1},
			},
			want: 878,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOddLengthSubarrays(tt.args.arr); got != tt.want {
				t.Errorf("sumOddLengthSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
