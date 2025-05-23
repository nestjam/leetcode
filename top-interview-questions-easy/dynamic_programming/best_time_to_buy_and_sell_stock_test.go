package dynamicprogramming

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			want: 1,
			args: args{
				prices: []int{1, 2},
			},
		},
		{
			want: 2,
			args: args{
				prices: []int{1, 2, 3},
			},
		},
		{
			want: 5,
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
		},
		{
			want: 0,
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
