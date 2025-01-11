package trie

import (
	"reflect"
	"testing"
)

func Test_stringMatching(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				words: []string{"a"},
			},
			want: []string{},
		},
		{
			args: args{
				words: []string{"foo", "foobar"},
			},
			want: []string{"foo"},
		},
		{
			args: args{
				words: []string{"mass", "as", "hero", "superhero"},
			},
			want: []string{"as", "hero"},
		},
		{
			args: args{
				words: []string{"leetcode", "cod", "code"},
			},
			want: []string{"cod", "code"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringMatching(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}
