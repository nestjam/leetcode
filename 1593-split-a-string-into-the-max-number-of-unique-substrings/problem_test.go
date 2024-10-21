package splitastringintothemaxnumberofuniquesubstrings

import "testing"

func Test_maxUniqueSplit(t *testing.T) {
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
				s: "a",
			},
			want: 1,
		},
		{
			args: args{
				s: "ab",
			},
			want: 2,
		},
		{
			args: args{
				s: "aa",
			},
			want: 1,
		},
		{
			args: args{
				s: "abc",
			},
			want: 3,
		},
		{
			args: args{
				s: "ababccc",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxUniqueSplit(tt.args.s); got != tt.want {
				t.Errorf("maxUniqueSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
