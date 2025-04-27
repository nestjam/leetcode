package reversewordsinastring

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
				s: "a good   example",
			},
			want: "example good a",
		},
		{
			args: args{
				s: "ab bc",
			},
			want: "bc ab",
		},
		{
			args: args{
				s: "ab",
			},
			want: "ab",
		},
		{
			args: args{
				s: " ab ",
			},
			want: "ab",
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
