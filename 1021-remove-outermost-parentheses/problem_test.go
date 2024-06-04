package removeoutermostparentheses

import "testing"

func Test_removeOuterParentheses(t *testing.T) {
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
				s: "()",
			},
			want: "",
		},
		{
			args: args{
				s: "()()",
			},
			want: "",
		},
		{
			args: args{
				s: "(())",
			},
			want: "()",
		},
		{
			args: args{
				s: "(()())(())",
			},
			want: "()()()",
		},
		{
			args: args{
				s: "(()())(())(()(()))",
			},
			want: "()()()()(())",
		},
		{
			args: args{
				s: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOuterParentheses(tt.args.s); got != tt.want {
				t.Errorf("removeOuterParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
