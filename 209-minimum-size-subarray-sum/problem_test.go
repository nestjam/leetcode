package minimumsizesubarraysum

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				target: 7,
				nums:   []int{2,3,1,2,4,3},
			},
			want: 2,
		},
		{
			args: args{
				target: 7,
				nums:   []int{2,3,1,2,4,3},
			},
			want: 2,
		},
		{
			args: args{
				target: 2,
				nums:   []int{0, 1, 1, 2},
			},
			want: 1,
		},
		{
			args: args{
				target: 2,
				nums:   []int{0, 1, 1, 0},
			},
			want: 2,
		},
		{
			args: args{
				target: 2,
				nums:   []int{0, 0},
			},
			want: 0,
		},
		{
			args: args{
				target: 2,
				nums:   []int{0, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
