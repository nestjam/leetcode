package shortestdistancetoacharacter

import (
	"reflect"
	"testing"
)

func Test_shortestToChar(t *testing.T) {
	type args struct {
		s string
		c byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				s: "c",
				c: 'c',
			},
			want: []int{0},
		},
		{
			args: args{
				s: "abc",
				c: 'c',
			},
			want: []int{2, 1, 0},
		},
		{
			args: args{
				s: "abc",
				c: 'a',
			},
			want: []int{0, 1, 2},
		},
		{
			args: args{
				s: "abc",
				c: 'b',
			},
			want: []int{1, 0, 1},
		},
		{
			args: args{
				s: "aba",
				c: 'a',
			},
			want: []int{0, 1, 0},
		},
		{
			args: args{
				s: "abcba",
				c: 'a',
			},
			want: []int{0, 1, 2, 1, 0},
		},
		{
			args: args{
				s: "loveleetcode",
				c: 'e',
			},
			want: []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestToChar(tt.args.s, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestToChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
