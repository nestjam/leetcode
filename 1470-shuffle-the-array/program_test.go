package shufflethearray

import (
	"reflect"
	"testing"
)

func Test_shuffle(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				n:    1,
				nums: []int{1, 2},
			},
			want: []int{1, 2},
		},
		{
			args: args{
				n:    2,
				nums: []int{1, 2, 3, 4},
			},
			want: []int{1, 3, 2, 4},
		},
		{
			args: args{
				n:    3,
				nums: []int{2, 5, 1, 3, 4, 7},
			},
			want: []int{2, 3, 5, 4, 1, 7},
		},
		{
			args: args{
				n:    4,
				nums: []int{1, 2, 3, 4, 4, 3, 2, 1},
			},
			want: []int{1, 4, 2, 3, 3, 2, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shuffle(tt.args.nums, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}
