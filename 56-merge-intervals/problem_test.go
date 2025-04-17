package mergeintervals

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				intervals: [][]int{{1, 2}, {0, 1}},
			},
			want: [][]int{{0, 2}},
		},
		{
			args: args{
				intervals: [][]int{{0, 1}, {2, 3}},
			},
			want: [][]int{{0, 1}, {2, 3}},
		},
		{
			args: args{
				intervals: [][]int{{0, 3}, {1, 2}},
			},
			want: [][]int{{0, 3}},
		},
		{
			args: args{
				intervals: [][]int{{0, 1}, {1, 2}},
			},
			want: [][]int{{0, 2}},
		},
		{
			args: args{
				intervals: [][]int{{0, 1}},
			},
			want: [][]int{{0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
