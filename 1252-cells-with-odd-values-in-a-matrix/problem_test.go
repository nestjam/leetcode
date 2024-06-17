package cellswithoddvaluesinamatrix

import "testing"

func Test_oddCells(t *testing.T) {
	type args struct {
		m       int
		n       int
		indices [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				m:       1,
				n:       1,
				indices: [][]int{{0, 0}},
			},
			want: 0,
		},
		{
			args: args{
				m:       2,
				n:       1,
				indices: [][]int{{0, 0}},
			},
			want: 1,
		},
		{
			args: args{
				m:       2,
				n:       2,
				indices: [][]int{{0, 0}, {0, 1}},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddCells(tt.args.m, tt.args.n, tt.args.indices); got != tt.want {
				t.Errorf("oddCells() = %v, want %v", got, tt.want)
			}
		})
	}
}
