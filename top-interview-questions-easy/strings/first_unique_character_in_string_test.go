package strings

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{ s: "a" },
			want: 0,
		},
		{
			args: args{ s: "aa" },
			want: -1,
		},
		{
			args: args{ s: "leetcode" },
			want: 0,
		},
		{
			args: args{ s: "loveleetcode" },
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
