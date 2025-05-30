package longestpalindromicsubstring

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			args: args{
				s: "caabaabbaa",
			},
			want: "aabbaa",
		},
		{
			args: args{
				s: "aaba",
			},
			want: "aba",
		},
		{
			args: args{
				s: "a",
			},
			want: "a",
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
