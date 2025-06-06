package howmanynumbersaresmallerthanthecurrentnumber

import (
	"reflect"
	"testing"
)

func Test_smallerNumbersThanCurrent(t *testing.T) {
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
				nums: []int{0, 0},
			},
			want: []int{0, 0},
		},
		{
			args: args{
				nums: []int{0, 1},
			},
			want: []int{0, 1},
		},
		{
			args: args{
				nums: []int{8, 1, 2, 2, 3},
			},
			want: []int{4, 0, 1, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallerNumbersThanCurrent(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallerNumbersThanCurrent() = %v, want %v", got, tt.want)
			}
		})
	}
}
