package kclosestpointstoorigin

import (
	"reflect"
	"testing"
)

func Test_kClosest(t *testing.T) {
	type args struct {
		points [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				k:      1,
				points: [][]int{{0, 0}},
			},
			want: [][]int{{0, 0}},
		},
		{
			args: args{
				k:      2,
				points: [][]int{{0, 0}, {0, 0}, {0, 0}},
			},
			want: [][]int{{0, 0}, {0, 0}},
		},
		{
			args: args{
				k:      1,
				points: [][]int{{1, 3}, {-2, 2}},
			},
			want: [][]int{{-2, 2}},
		},
		{
			args: args{
				k:      2,
				points: [][]int{{3, 3}, {-2, 4}, {5, 1}},
			},
			want: [][]int{{3, 3}, {-2, 4}},
		},
		{
			args: args{
				k:      1,
				points: [][]int{{10000, 10000}},
			},
			want: [][]int{{10000, 10000}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kClosest(tt.args.points, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
