package selfdividingnumbers

import (
	"reflect"
	"testing"
)

func Test_selfDividingNumbers(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				left:  1,
				right: 1,
			},
			want: []int{1},
		},
		{
			args: args{
				left:  1,
				right: 2,
			},
			want: []int{1, 2},
		},
		{
			args: args{
				left:  11,
				right: 13,
			},
			want: []int{11, 12},
		},
		{
			args: args{
				left:  9,
				right: 13,
			},
			want: []int{9, 11, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selfDividingNumbers(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selfDividingNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
