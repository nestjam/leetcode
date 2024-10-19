package maximumswap

import "testing"

func Test_maximumSwap(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				num: 1,
			},
			want: 1,
		},
		{
			args: args{
				num: 12,
			},
			want: 21,
		},
		{
			args: args{
				num: 10,
			},
			want: 10,
		},
		{
			args: args{
				num: 112,
			},
			want: 211,
		},
		{
			args: args{
				num: 312,
			},
			want: 321,
		},
		{
			args: args{
				num: 9973,
			},
			want: 9973,
		},
		{
			args: args{
				num: 1993,
			},
			want: 9913,
		},
		{
			args: args{
				num: 98368,
			},
			want: 98863,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSwap(tt.args.num); got != tt.want {
				t.Errorf("maximumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
