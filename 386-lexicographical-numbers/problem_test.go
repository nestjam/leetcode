package lexicographicalnumbers

import (
	"reflect"
	"testing"
)

func Test_lexicalOrder(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				n: 1,
			},
			want: []int{1},
		},
		{
			args: args{
				n: 9,
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			args: args{
				n: 10,
			},
			want: []int{1, 10, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			args: args{
				n: 13,
			},
			want: []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			args: args{
				n: 21,
			},
			want: []int{1, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 2, 20, 21, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lexicalOrder(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexicalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
