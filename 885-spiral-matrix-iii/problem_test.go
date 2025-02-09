package spiralmatrixiii

import (
	"reflect"
	"testing"
)

func Test_spiralMatrixIII(t *testing.T) {
	type args struct {
		rows   int
		cols   int
		rStart int
		cStart int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				rows:   1,
				cols:   1,
				rStart: 0,
				cStart: 0,
			},
			want: [][]int{{0, 0}},
		},
		{
			args: args{
				rows:   1,
				cols:   4,
				rStart: 0,
				cStart: 0,
			},
			want: [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		},
		{
			args: args{
				rows:   5,
				cols:   6,
				rStart: 1,
				cStart: 4,
			},
			want: [][]int{{1, 4}, {1, 5}, {2, 5}, {2, 4}, {2, 3}, {1, 3}, {0, 3}, {0, 4}, {0, 5}, {3, 5}, {3, 4}, {3, 3}, {3, 2}, {2, 2}, {1, 2}, {0, 2}, {4, 5}, {4, 4}, {4, 3}, {4, 2}, {4, 1}, {3, 1}, {2, 1}, {1, 1}, {0, 1}, {4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralMatrixIII(tt.args.rows, tt.args.cols, tt.args.rStart, tt.args.cStart); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralMatrixIII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNext(t *testing.T) {
	type args struct {
		moveCount int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				moveCount: 0,
			},
			want: []int{0, 0},
		},
		{
			args: args{
				moveCount: 1,
			},
			want: []int{0, 1},
		},
		{
			args: args{
				moveCount: 2,
			},
			want: []int{1, 1},
		},
		{
			args: args{
				moveCount: 3,
			},
			want: []int{1, 0},
		},
		{
			args: args{
				moveCount: 4,
			},
			want: []int{1, -1},
		},
		{
			args: args{
				moveCount: 5,
			},
			want: []int{0, -1},
		},
		{
			args: args{
				moveCount: 6,
			},
			want: []int{-1, -1},
		},
		{
			args: args{
				moveCount: 7,
			},
			want: []int{-1, 0},
		},
		{
			args: args{
				moveCount: 9,
			},
			want: []int{-1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iterator := newSpiralIterator(0, 0)
			for i := 0; i < tt.args.moveCount; i++ {
				iterator.moveNext()
			}
			if got := iterator.current(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
