package reversewordsinastringiii

import "testing"

func Test_reverseWords(t *testing.T) {
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
				s: "abc",
			},
			want: "cba",
		},
		{
			args: args{
				s: "",
			},
			want: "",
		},
		{
			args: args{
				s: "hello world!",
			},
			want: "olleh !dlrow",
		},
		{
			args: args{
				s: "Let's take LeetCode contest",
			},
			want: "s'teL ekat edoCteeL tsetnoc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
