package intersectionoftwoarrays

import (
	"reflect"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums1: []int{1},
				nums2: []int{1},
			},
			want: []int{1},
		},
		{
			args: args{
				nums1: []int{0},
				nums2: []int{1},
			},
			want: []int{},
		},
		{
			args: args{
				nums1: []int{1},
				nums2: []int{0},
			},
			want: []int{},
		},
		{
			args: args{
				nums1: []int{1},
				nums2: []int{0, 1},
			},
			want: []int{1},
		},
		{
			args: args{
				nums1: []int{1, 0},
				nums2: []int{1},
			},
			want: []int{1},
		},
		{
			args: args{
				nums1: []int{4, 9, 5},
				nums2: []int{9, 4, 9, 8, 4},
			},
			want: []int{4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
