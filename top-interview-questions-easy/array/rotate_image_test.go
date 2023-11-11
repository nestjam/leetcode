package array

import (
	"reflect"
	"testing"
)

func Test_rotateImage(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		want [][]int
		args args
	}{
		{
			want: [][]int{{3, 1}, {4, 2}},
			args: args{
				matrix: [][]int{{1, 2}, {3, 4}},
			},
		},
		{
			want: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
			args: args{
				matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
		},
		{
			want: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
			args: args{
				matrix: [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotateImage(tt.args.matrix)
			got := tt.args.matrix
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
