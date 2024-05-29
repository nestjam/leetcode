package flippinganimage

import (
	"reflect"
	"testing"
)

func Test_flipAndInvertImage(t *testing.T) {
	type args struct {
		image [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				image: [][]int{{1}},
			},
			want: [][]int{{0}},
		},
		{
			args: args{
				image: [][]int{{0, 1, 0}},
			},
			want: [][]int{{1, 0, 1}},
		},
		{
			args: args{
				image: [][]int{{1, 1}},
			},
			want: [][]int{{0, 0}},
		},
		{
			args: args{
				image: [][]int{{0, 1, 0}, {1, 0, 1}},
			},
			want: [][]int{{1, 0, 1}, {0, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipAndInvertImage(tt.args.image); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipAndInvertImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
