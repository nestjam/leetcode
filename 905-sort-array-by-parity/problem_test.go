package sortarraybyparity

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParity(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{1, 2},
			},
			want: []int{2, 1},
		},
		{
			args: args{
				nums: []int{4, 2},
			},
			want: []int{4, 2},
		},
		{
			args: args{
				nums: []int{3, 1, 2, 4},
			},
			want: []int{2, 4, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParity(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}
