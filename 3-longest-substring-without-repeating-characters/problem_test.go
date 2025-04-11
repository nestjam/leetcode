package longestsubstringwithoutrepeatingcharacters

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
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
				"aabaab!bb",
			},
			want: 3,
		},
		{
			args: args{
				"aab",
			},
			want: 2,
		},
		{
			args: args{
				"abcabcbb",
			},
			want: 3,
		},
		{
			args: args{
				"ab",
			},
			want: 2,
		},
		{
			args: args{
				"aa",
			},
			want: 1,
		},
		{
			args: args{
				"",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
