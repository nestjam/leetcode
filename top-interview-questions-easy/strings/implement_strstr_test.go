package strings

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			want: 0,
			args: args{
				haystack: "abc",
				needle: "a",
			},
		},
		{
			want: 1,
			args: args{
				haystack: "abc",
				needle: "b",
			},
		},
		{
			want: 2,
			args: args{
				haystack: "abac",
				needle: "ac",
			},
		},
		{
			want: -1,
			args: args{
				haystack: "abac",
				needle: "df",
			},
		},
		{
			want: 4,
			args: args{
				haystack: "mississippi",
				needle: "issip",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
