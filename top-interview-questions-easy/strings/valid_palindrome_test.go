package strings

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{s: " "},
			want: true,
		},
		{
			args: args{s: "ab"},
			want: false,
		},
		{
			args: args{s: "aA"},
			want: true,
		},
		{
			args: args{s: "aA"},
			want: true,
		},
		{
			args: args{s: " aA"},
			want: true,
		},
		{
			args: args{s: "11"},
			want: true,
		},
		{
			args: args{s: "12"},
			want: false,
		},
		{
			args: args{s: "a "},
			want: true,
		},
		{
			args: args{s: "a _a"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
