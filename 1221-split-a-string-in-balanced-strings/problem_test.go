package splitastringinbalancedstrings

import "testing"

func Test_balancedStringSplit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s: "LR",
			},
			want: 1,
		},
		{
			args: args{
				s: "LRLR",
			},
			want: 2,
		},
		{
			args: args{
				s: "RL",
			},
			want: 1,
		},
		{
			args: args{
				s: "RLRRLLRLRL",
			},
			want: 4,
		},
		{
			args: args{
				s: "RLRRRLLRLL",
			},
			want: 2,
		},
		{
			args: args{
				s: "LLLLRRRR",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balancedStringSplit(tt.args.s); got != tt.want {
				t.Errorf("balancedStringSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
