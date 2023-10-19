package dynamicprogramming

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			want: 0,
			args: args{
				nums: []int { },
			},
		},
		{
			want: 1,
			args: args{
				nums: []int { 1 },
			},
		},
		{
			want: 4,
			args: args{
				nums: []int { 1,2,3,1 },
			},
		},
		{
			want: 12,
			args: args{
				nums: []int { 2,7,9,3,1 },
			},
		},
		{
			want: 4,
			args: args{
				nums: []int { 2,1,1,2 },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
