package gameoflife

import (
	"reflect"
	"testing"
)

func Test_gameOfLife(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				board: [][]int{{1, 1}, {1, 0}},
			},
			want: [][]int{{1, 1}, {1, 1}},
		},
		{
			args: args{
				board: [][]int{{1, 0}, {0, 0}},
			},
			want: [][]int{{0, 0}, {0, 0}},
		},
		{
			args: args{
				board: [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}},
			},
			want: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gameOfLife(tt.args.board); !reflect.DeepEqual(tt.args.board, tt.want) {
				t.Errorf("intersection() = %v, want %v", tt.args.board, tt.want)
			}
		})
	}
}
