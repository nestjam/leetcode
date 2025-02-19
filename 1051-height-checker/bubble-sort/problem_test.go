package bubblesort

import "testing"

func Test_heightChecker(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				heights: []int{1, 2},
			},
			want: 0,
		},
		{
			args: args{
				heights: []int{2, 1},
			},
			want: 2,
		},
		{
			args: args{
				heights: []int{1, 1, 4, 2, 1, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heightChecker(tt.args.heights); got != tt.want {
				t.Errorf("heightChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
