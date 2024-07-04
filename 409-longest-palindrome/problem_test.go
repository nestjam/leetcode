package longestpalindrome

import "testing"

func Test_longestPalindrome(t *testing.T) {
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
				s: "aa",
			},
			want: 2,
		},
		{
			args: args{
				s: "ba",
			},
			want: 1,
		},
		{
			args: args{
				s: "abccccdd",
			},
			want: 7,
		},
		{
			args: args{
				s: "abbbccccc",
			},
			want: 7,
		},
		{
			args: args{
				s: "zeusnilemacaronimaisanitratetartinasiaminoracamelinsuez",
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
