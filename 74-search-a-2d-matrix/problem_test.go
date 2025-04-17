package searcha2dmatrix

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				matrix: [][]int{{1, 3}, {4, 5}},
				target: 2,
			},
			want: false,
		},
		{
			args: args{
				matrix: [][]int{{1, 2}, {3, 4}},
				target: 20,
			},
			want: false,
		},
		{
			args: args{
				matrix: [][]int{{1, 2}, {3, 4}},
				target: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
