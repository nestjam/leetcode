package strings

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			want: "flow",
			args: args{
				strs: []string { "flower", "flow" },
			},
		},
		{
			want: "",
			args: args{
				strs: []string { "a", "b" },
			},
		},
		{
			want: "",
			args: args{
				strs: []string { },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
