package robotreturntoorigin

import "testing"

func Test_judgeCircle(t *testing.T) {
	type args struct {
		moves string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				moves: "UD",
			},
			want: true,
		},
		{
			args: args{
				moves: "L",
			},
			want: false,
		},
		{
			args: args{
				moves: "R",
			},
			want: false,
		},
		{
			args: args{
				moves: "U",
			},
			want: false,
		},
		{
			args: args{
				moves: "D",
			},
			want: false,
		},
		{
			args: args{
				moves: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeCircle(tt.args.moves); got != tt.want {
				t.Errorf("judgeCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}
