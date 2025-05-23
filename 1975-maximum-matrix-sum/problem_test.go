package maximummatrixsum

import "testing"

func Test_maxMatrixSum(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				matrix: [][]int{{1, 1}, {1, 1}},
			},
			want: 4,
		},
		{
			args: args{
				matrix: [][]int{{-1, -1}, {1, 1}},
			},
			want: 4,
		},
		{
			args: args{
				matrix: [][]int{{1, 1}, {-1, -1}},
			},
			want: 4,
		},
		{
			args: args{
				matrix: [][]int{{-1, 1}, {-1, 1}},
			},
			want: 4,
		},
		{
			args: args{
				matrix: [][]int{{1, -1}, {1, -1}},
			},
			want: 4,
		},
		{
			args: args{
				matrix: [][]int{{-1, 1}, {1, -1}},
			},
			want: 4,
		},
		{
			args: args{
				matrix: [][]int{{-2, 1}, {1, 1}},
			},
			want: 3,
		},
		{
			args: args{
				matrix: [][]int{{-1, 2}, {2, 1}},
			},
			want: 4,
		},
		{
			args: args{
				matrix: [][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxMatrixSum(tt.args.matrix); got != tt.want {
				t.Errorf("maxMatrixSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
