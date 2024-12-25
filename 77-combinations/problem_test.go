package combinations

import (
	"reflect"
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				n: 2,
				k: 1,
			},
			want: [][]int{{1}, {2}},
		},
		{
			args: args{
				n: 2,
				k: 2,
			},
			want: [][]int{{2, 1}},
		},
		{
			args: args{
				n: 4,
				k: 2,
			},
			want: [][]int{{2, 1}, {3, 1}, {4, 1}, {3, 2}, {4, 2}, {4, 3}},
		},
		{
			args: args{
				n: 3,
				k: 3,
			},
			want: [][]int{{3, 2, 1}},
		},
		{
			args: args{
				n: 4,
				k: 3,
			},
			want: [][]int{{3, 2, 1}, {4, 2, 1}, {4, 3, 1}, {4, 3, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
